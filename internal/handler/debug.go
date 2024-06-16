package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserId   uuid.UUID `db:"user_id"`
	UserName string    `db:"user_name"`
	RoomId   string    `db:"room_id"`
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err).SetInternal(err)
	}
	return c.JSON(http.StatusOK, users)
}
