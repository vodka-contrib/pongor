package main

import (
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())
	r := pongor.GetRenderer()
	serv.SetRenderer(r)
	serv.Static("/static", "./static")
	serv.Get("/", func(ctx *echo.Context) error {
		ctx.Render(200, "index.html", map[string]interface{}{
			"title": "你好，世界",
		})
		return nil
	})

	serv.Run("127.0.0.1:9000")
}
