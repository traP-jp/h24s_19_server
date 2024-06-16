package repository

import (
	"context" // Add the context package

	"github.com/google/uuid"
)

type Room struct {
	RoomId   uuid.UUID `db:"room_id" json:"roomId"`
	RoomName string    `db:"room_name" json:"roomName"`
	IsPublic bool      `db:"is_public" json:"isPublic"`
}

type RoomRequest struct {
	RoomName string `json:"room_name"`
	IsPublic bool   `json:"is_public"`
	Password string `json:"password"`
}

func (r *Repository) GetRooms(ctx context.Context) ([]Room, error) {
	var rooms []Room
	err := r.db.Select(&rooms, "SELECT * FROM rooms")
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *Repository) CreateRoom(ctx context.Context, params RoomRequest) (Room, error) {
	roomId, err := uuid.NewUUID()
	if err != nil {
		return Room{}, err
	}
	_, err = r.db.Exec(
		"INSERT INTO rooms (room_id, room_name, is_public) VALUES (?, ?, ?)",
		roomId,
		params.RoomName,
		params.IsPublic,
	)
	if err != nil {
		return Room{}, err
	}
	room := Room{
		RoomId:   roomId,
		RoomName: params.RoomName,
		IsPublic: params.IsPublic,
	}
	return room, nil
}
