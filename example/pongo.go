package main

import (
	"github.com/vodka-contrib/pongor"
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/middleware"
)

func main() {
	v := vodka.New()
	v.Use(middleware.Logger())
	v.Use(middleware.Recover())
	r := pongor.Renderor()
	v.SetRenderer(r)
	v.Static("/static", "./static")
	v.Get("/", func(ctx *vodka.Context) error {
		ctx.Render(200, "index.html", map[string]interface{}{
			"title": "你好，世界",
		})
		return nil
	})

	v.Run("127.0.0.1:9000")
}
