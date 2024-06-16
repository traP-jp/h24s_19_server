package streamer

import (
	"h24s_19/internal/pkg/util"

	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type postWordArgs struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

type postWordResponse struct {
	Type       string `json:"type"`
	UserName   string `json:"userName"`
	WordId     int    `json:"wordId"`
	Word       string `json:"word"`
	Reading    string `json:"reading"`
	BasicScore int    `json:"basicScore"`
}

type rejectedPostWord struct {
	Type    string `json:"type"`
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

var NotMatchShiritoriError = errors.New("ルール違反")

func addWord(db *sqlx.DB, roomId string, userId string, word string, reading string, basic_score int) (int, error) {
	lastReading, err := getLastWordReading(db, roomId)
	if err != nil {
		fmt.Println("failed to get last word:", err)
		return 0, err
	}
	fmt.Println("lastReading, reading, Check: ", lastReading, reading, util.SiritoriCheck(lastReading, reading))
	if lastReading != "" && !util.SiritoriCheck(lastReading, reading) {
		return 0, NotMatchShiritoriError
	}

	fmt.Println("word: %s, reading: %s, basic_score: %d", word, reading, basic_score)
	res, err := db.Exec("INSERT INTO words (room_id, user_id, word, reading, basic_score) VALUES (?, ?, ?, ?)", roomId, userId, word, reading, basic_score)
	if err != nil {
		fmt.Println("failed to insert word:", err)
		return 0, err
	}
	wordId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("failed to get last insert id: %v")
		return 0, err
	}

	return int(wordId), nil
}

func getLastWordReading(db *sqlx.DB, roomId string) (string, error) {
	var reading string
	err := db.Get(&reading, "SELECT reading FROM words WHERE room_id = ? ORDER BY word_id DESC LIMIT 1", roomId)
	if err != nil {
		return "", nil
	}
	return reading, nil
}

func (s *Streamer) handleInit(db *sqlx.DB, roomId string, clientID uuid.UUID) error {
	var words []struct {
		WordId     int    `db:"word_id"`
		Word       string `db:"word"`
		Reading    string `db:"reading"`
		UserName   string `db:"user_name"`
		BasicScore int    `db:"basic_score"`
	}
	err := db.Select(&words, "SELECT `word_id`, `word`, `reading`, `user_name`, `basic_score` FROM `words` WHERE `room_id` = ? ORDER BY `word_id` LIMIT 10", roomId)
	if err != nil {
		fmt.Println("failed to get words:", err)
		return err
	}

	for _, word := range words {
		message := postWordResponse{
			Type:       "posted_word",
			UserName:   word.UserName,
			WordId:     word.WordId,
			Word:       word.Word,
			Reading:    word.Reading,
			BasicScore: word.BasicScore,
		}
		messageBytes, err := json.Marshal(message)
		if err != nil {
			fmt.Println("failed to marshal message: ", err)
			return err
		}
		s.sendTo(string(messageBytes), func(c *client) bool {
			return c.id == clientID
		})
	}
	return nil
}

func (s *Streamer) handlePostWord(db *sqlx.DB, roomId string, clientID uuid.UUID, args postWordArgs) error {
	lastReading, err := getLastWordReading(db, roomId)
	if err != nil {
		fmt.Println("failed to get last word:", err)
		return err
	}
	userId := s.clients[clientID].name
	roomIdUuid := uuid.FromStringOrNil(roomId)
	rune_count, err := s.repo.GetRuneCount(roomIdUuid)
	basicScore := util.GetScore(lastReading, args.Reading, rune_count)
	wordId, err := addWord(db, roomId, userId, args.Word, args.Reading, basicScore)
	if err != nil {
		if err == NotMatchShiritoriError {
			message := rejectedPostWord{
				Type:    "post_word_rejected",
				Word:    args.Word,
				Reading: args.Reading,
			}
			messageBytes, err := json.Marshal(message)
			if err != nil {
				fmt.Println("failed to marshal message: ", err)
				return err
			}
			s.sendTo(string(messageBytes), func(c *client) bool {
				return c.id == clientID
			})
		}

		return err
	}
	message := postWordResponse{
		Type:       "posted_word",
		UserName:   s.clients[clientID].name,
		WordId:     wordId,
		Word:       args.Word,
		Reading:    args.Reading,
		BasicScore: basicScore,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("failed to marshal message: ", err)
		return err
	}
	fmt.Println("message: ", string(messageBytes))

	s.sendToRoom(roomId, string(messageBytes))

	return nil
}
