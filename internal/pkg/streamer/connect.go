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
	fmt.Println("connectWS")
	roomID, err := c.Cookie("roomID")
	if err != nil {
		fmt.Println("room cookie not set")
		return ErrorRoomNotFound
	}
	fmt.Println("roomID: ", roomID.Value)
	_, err = s.repo.GetRoom(c.Request().Context(), roomID.Value)
	if err != nil {
		fmt.Println("room not found")
		return ErrorRoomNotFound
	}

	fmt.Println("部屋取れた")

	userIdInCookie, err := c.Cookie("userId")
	if err != nil {
		fmt.Println("user cookie not set")
		return ErrorUserCookieNotSet
	}
	userId := userIdInCookie.Value

	fmt.Println("userId: ", userId)

	user, err := s.repo.GetUser(c.Request().Context(), userId)
	fmt.Println("user:", user, "err:", err)
	if err != nil {
		return ErrorUserNotFound
	}

	fmt.Println("2")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Println("3")

	client := newClient(roomID.Value, conn, s.receiver)

	fmt.Println("4")

	s.clients[client.id] = client
	s.clients[client.id].name = user.UserName

	fmt.Println("5")

	go client.listen()
	go client.send()

	fmt.Println("6")

	<-client.closer

	fmt.Println("7")

	delete(s.clients, client.id)

	fmt.Println("8")

	return c.NoContent(http.StatusOK)
}
