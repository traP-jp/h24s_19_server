package handler

import (
	"github.com/labstack/echo/v4"
	"h24s_19/internal/pkg/streamer"
	"h24s_19/internal/repository"
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

	// ws API
	api.GET("/ws/:roomID", h.streamer.ConnectWS)
}
