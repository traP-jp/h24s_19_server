package handler

import (
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
	RoomName string `json:"room_name"`
	IsPublic bool   `json:"is_public"`
	Password string `json:"password"`
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

