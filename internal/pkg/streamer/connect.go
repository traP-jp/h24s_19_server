package streamer

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

var ErrorRoomNotFound = echo.NewHTTPError(http.StatusNotFound, "room not found")
var ErrorUserCookieNotSet = echo.NewHTTPError(http.StatusUnauthorized, "user cookie not set")
var ErrorUserNotFound = echo.NewHTTPError(http.StatusNotFound, "user not found")

func (s *Streamer) ConnectWS(c echo.Context) error {
	roomID := c.Param("roomID")
	_, err := s.repo.GetRoom(c.Request().Context(), roomID)
	if err != nil {
		return ErrorRoomNotFound
	}

	userIdInCookie, err := c.Cookie("userId")
	if err != nil {
		return ErrorUserCookieNotSet
	}
	userId := userIdInCookie.Value

	fmt.Println("userId: ", userId)

	user, error := s.repo.GetUser(c.Request().Context(), userId)
	if error != nil {
		return ErrorUserNotFound
	}

	fmt.Println("2")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Println("3")

	client := newClient(roomID, conn, s.receiver)

	s.clients[client.id] = client
	s.clients[client.id].name = user.UserName

	go client.listen()
	go client.send()

	<-client.closer

	delete(s.clients, client.id)

	return c.NoContent(http.StatusOK)
}
