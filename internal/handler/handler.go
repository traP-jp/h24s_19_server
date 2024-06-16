package handler

import (
	"h24s_19/internal/pkg/streamer"
	"h24s_19/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo     *repository.Repository
	streamer *streamer.Streamer
}

func New(repo *repository.Repository, streamer *streamer.Streamer) *Handler {
	return &Handler{
		repo:     repo,
		streamer: streamer,
	}
}

func (h *Handler) SetupRoutes(api *echo.Group) {

	api.GET("/rooms", h.GetRooms)

	api.POST("/room", h.CreateRoom)

	api.POST("/room/:roomId/enter", h.EnterRoom)

	// ws API
	api.GET("/ws/:roomID", h.streamer.ConnectWS)
}
