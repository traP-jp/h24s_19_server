package streamer

import (
	"h24s_19/internal/repository"
	"log"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type receiveData struct {
	clientID uuid.UUID
	roomID   string
	payload  []byte
}

type Streamer struct {
	clients  map[uuid.UUID]*client
	receiver chan receiveData
	repo     *repository.Repository
}

func NewStreamer(r *repository.Repository) *Streamer {
	return &Streamer{
		clients:  make(map[uuid.UUID]*client),
		receiver: make(chan receiveData),
		repo:     r,
	}
}

func (s *Streamer) Listen(db *sqlx.DB) {
	for {
		data := <-s.receiver

		go func() {
			err := s.handleWebSocket(db, data)
			if err != nil {
				log.Printf("failed to handle websocket: %v", err)
			}
		}()
	}
}

func (s *Streamer) sendToRoom(roomID, msg string) {
	log.Printf("send to room: %s, msg: %s", roomID, msg)
	for _, c := range s.clients {
		if c.roomID == roomID {
			c.sender <- msg
		}
	}
}

func (s *Streamer) sendTo(msg string, cond func(c *client) bool) error {
	for _, c := range s.clients {
		if cond(c) {
			c.sender <- msg
		}
	}
	return nil
}
