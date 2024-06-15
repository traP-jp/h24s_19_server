package streamer

import (
	"encoding/json"
	"fmt"
	"log"
)

type payload struct {
	Type string          `json:"type,omitempty"`
	Args json.RawMessage `json:"args,omitempty"`
}

type postWordArgs struct {
	Word string `json:"word"`
}

func (s *Streamer) handleWebSocket(data receiveData) error {
	var req payload
	err := json.Unmarshal(data.payload, &req)
	if err != nil {
		return err
	}

	switch req.Type {
	case "post_word":
		var args postWordArgs
		err = json.Unmarshal(req.Args, &args)
		if err != nil {
			return err
		}
		s.sendToRoom(data.roomID, args.Word)
	default:
		log.Printf("unknown type: %s", req.Type)
		return fmt.Errorf("unknown type: %s", req.Type)
	}

	return nil
}
