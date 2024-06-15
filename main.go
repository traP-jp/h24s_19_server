package main

import (
	"fmt"
	"h24s_19/internal/pkg/config"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

func main() {
	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", func(c echo.Context) error {
		// HTTPステータスコードは200番で、文字列「Hello, World.」をクライアントに返す
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}

	// 部屋一覧取得
	e.GET("/api/rooms", func(c echo.Context) error {
		var rooms []Room
		err := db.Select(&rooms, "SELECT * FROM rooms")
		if err != nil {
			e.Logger.Fatal(err)
			return err
		}
		return c.JSON(http.StatusOK, rooms)
	})

	// 部屋作成
	e.POST("/api/room", func(c echo.Context) error {
		data := &RoomRequest{}
		if err := c.Bind(data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
		}
		roomId, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		_, err = db.Exec(
			"INSERT INTO rooms (room_id, room_name, is_public) VALUES (?, ?, ?)",
			roomId,
			data.RoomName,
			data.IsPublic,
		)
		if err != nil {
			return err
		}
		room := Room{
			RoomId:   roomId,
			RoomName: data.RoomName,
			IsPublic: data.IsPublic,
		}
		return c.JSON(http.StatusOK, room)
	})

	defer db.Close()

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}
