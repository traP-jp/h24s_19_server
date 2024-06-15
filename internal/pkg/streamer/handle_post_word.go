package streamer

import (
	"encoding/json"
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

func (s *Streamer) handlePostWord(roomId string, args postWordArgs) error {
	// wordId, err := ..AddWord(roomId, args.Word, args.Reading)
	// if err != nil {
	// 	if err == NotMatchShiritoriError {
	// 		return s.send(rejectedPostWord{
	// 			Type: "post_word_rejected",
	// 			Word: args.Word,
	// 			Reading: args.Reading,
	// 		})
	// 	}
	// 	return err
	// }
	message := postWordResponse{
		Type:       "posted_word",
		WordId:     0, // wordId,
		Word:       args.Word,
		Reading:    args.Reading,
		BasicScore: 0, // GetBasicScore(args.Word),
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	s.sendToRoom(roomId, string(messageBytes))

	return nil
}
