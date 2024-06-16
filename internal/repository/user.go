package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID `db:"user_id"`
	UserName string    `db:"user_name"`
	RoomId   string    `db:"room_id"`
}

type CreateUserRequest struct {
	UserName     string `json:"userName"`
	RoomId       string `json:"roomId"`
	RoomPassword string `json:"password"`
}

var NotEnteredRoomError = errors.New("ルームに入っていません、まず GET /api/room/:roomId/enter")

func (r *Repository) GetUser(ctx context.Context, userId string) (User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM `users` WHERE `user_id` = ?", userId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return User{}, NotEnteredRoomError
		}
		fmt.Println("failed to get user:", err)
		return User{}, err
	}
	fmt.Println("user: ", user)
	return user, nil
}

func (r *Repository) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := r.db.Select(&users, "SELECT * FROM `users`")
	if err != nil {
		return nil, err
	}
	return users, nil
}

var ErrorRoomNotFound = errors.New("ルームが見つかりません")
var ErrorNotMatchRoomPassword = errors.New("パスワードが違います")

func (r *Repository) CreateUser(ctx context.Context, params CreateUserRequest) (User, error) {
	_, err := r.GetRoom(ctx, params.RoomId)
	if err != nil {
		return User{}, ErrorRoomNotFound
	}

	if false { // check room password
		return User{}, ErrorNotMatchRoomPassword
	}

	userId, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}
	_, err = r.db.Exec(
		"INSERT INTO `users` (`user_id`, `user_name`, `room_id`) VALUES (?, ?, ?)",
		userId,
		params.UserName,
		params.RoomId,
	)
	if err != nil {
		fmt.Println("failed to insert user:", err)
		return User{}, err
	}
	user := User{
		UserId:   userId,
		UserName: params.UserName,
	}
	return user, nil
}
