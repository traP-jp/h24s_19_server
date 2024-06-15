package main

import (
	"net/http"
	"h24s_19/internal/pkg/config"
	

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Room struct {
	roomId string `db:"room_id"`
	roomName string `db:"room_name"`
	isPublic bool `db:"is_public"`
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
	e.GET("/api/rooms", func(c echo.Context) error {
		var rooms []Room
		err := db.Select(&rooms, "SELECT * FROM rooms")
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, rooms)
	})

	defer db.Close()

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}
