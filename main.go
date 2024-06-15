package main

import (
	"h24s_19/internal/handler"
	"h24s_19/internal/pkg/config"
	"h24s_19/internal/pkg/streamer"
	"h24s_19/internal/repository"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	s := streamer.NewStreamer()
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
	defer db.Close()

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo, s)
	v1API := e.Group("/api")
	h.SetupRoutes(v1API)

	go s.Listen()

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}
