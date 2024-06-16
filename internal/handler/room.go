package handler

import (
	"errors"
	"h24s_19/internal/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Room struct {
	RoomId   uuid.UUID `db:"room_id"`
	RoomName string    `db:"room_name"`
	IsPublic bool      `db:"is_public"`
}

type RoomRequest struct {
	IsPublic bool   `json:"isPublic"`
	RoomName string `json:"roomName"`
	Password string `json:"password"`
}

type EnterRoomRequest struct {
	UserName     string `json:"userName"`
	RoomPassword string `json:"password"`
}

type EnterRoomResponse struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func (h *Handler) GetRooms(c echo.Context) error {
	rooms, err := h.repo.GetRooms(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	return c.JSON(http.StatusOK, rooms)
}

func (h *Handler) CreateRoom(c echo.Context) error {
	var params RoomRequest
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}
	room, err := h.repo.CreateRoom(c.Request().Context(), repository.RoomRequest{params.RoomName, params.IsPublic, params.Password})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	return c.JSON(http.StatusCreated, room)
}

var ErrorRoomNotFound = errors.New("ルームが見つかりません")
var ErrorNotMatchRoomPassword = errors.New("パスワードが違います")

func (h *Handler) EnterRoom(c echo.Context) error {
	roomId := c.Param("roomId")
	if roomId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorRoomNotFound)
	}

	var params EnterRoomRequest
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}

	if false { // check room password
		return echo.NewHTTPError(http.StatusUnauthorized, ErrorNotMatchRoomPassword)
	}

	_, err := h.repo.GetRoom(c.Request().Context(), roomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrorRoomNotFound)
	}

	user, err := h.repo.CreateUser(c.Request().Context(), repository.CreateUserRequest{UserName: params.UserName, RoomId: roomId, RoomPassword: params.RoomPassword})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	// set cookie
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = user.UserId.String()
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, EnterRoomResponse{UserId: user.UserId.String(), UserName: user.UserName})
}
