package streamer

import (
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type postWordArgs struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

type postWordResponse struct {
	Type       string `json:type,omitempty`
	WordId     int    `json:wordId,omitempty`
	Word       string `json:word,omitempty`
	Reading    string `json:reading,omitempty`
	BasicScore int    `json:basicScore,omitempty`
}

type rejectedPostWord struct {
	Type    string `json:type,omitempty`
	Word    string `json:word,omitempty`
	Reading string `json:reading,omitempty`
}

var NotMatchShiritoriError = errors.New("ルール違反")

func addWord(db *sqlx.DB, roomId string, word string, reading string, basic_score int) (int, error) {
	if false { // ..TestShiritori(word)
		return 0, NotMatchShiritoriError
	}
	var wordId int
	db.QueryRow("INSERT INTO words (room_id, word, reading, basic_score) VALUES (?, ?, ?, ?) RETURNING word_id", roomId, word, reading, basic_score).Scan(&wordId)

	return wordId, nil
}

func (s *Streamer) handlePostWord(db *sqlx.DB, roomId string, clientID uuid.UUID, args postWordArgs) error {
	basicScore := 0 // ..GetBasicScore(args.Word)
	wordId, err := addWord(db, roomId, args.Word, args.Reading, basicScore)
	if err != nil {
		if err == NotMatchShiritoriError {
			message := rejectedPostWord{
				Type:    "post_word_rejected",
				Word:    args.Word,
				Reading: args.Reading,
			}
			messageBytes, err := json.Marshal(message)
			if err != nil {
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
		WordId:     wordId,
		Word:       args.Word,
		Reading:    args.Reading,
		BasicScore: basicScore,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	s.sendToRoom(roomId, string(messageBytes))

	return nil
}
