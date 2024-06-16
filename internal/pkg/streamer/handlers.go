package streamer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type payload struct {
	Type string          `json:"type,omitempty"`
	Args json.RawMessage `json:"args,omitempty"`
}

func (s *Streamer) handleWebSocket(db *sqlx.DB, data receiveData) error {
	var req payload
	err := json.Unmarshal(data.payload, &req)
	if err != nil {
		return err
	}

	switch req.Type {
	case "postWord":
		var args postWordArgs
		err = json.Unmarshal(req.Args, &args)
		if err != nil {
			return err
		}
		s.handlePostWord(db, data.roomID, data.clientID, args)
	default:
		log.Printf("unknown type: %s", req.Type)
		return fmt.Errorf("unknown type: %s", req.Type)
	}

	return nil
}
