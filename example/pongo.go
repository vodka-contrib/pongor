package main

import (
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/fasthttp"
	"github.com/insionng/vodka/middleware"
	"github.com/vodka-contrib/pongor"
)

func main() {
	v := vodka.New()
	v.Use(middleware.Logger())
	v.Use(middleware.Recover())
	r := pongor.Renderor()
	v.SetRenderer(r)
	v.Static("/static", "static")
	v.Get("/", func(ctx vodka.Context) error {
		var data = make(map[string]interface{})
		data["name"] = "Insion Ng"
		ctx.SetStore(data)

		ctx.SetStore(map[string]interface{}{
			"title": "你好，世界",
			"oh":    "no",
		})
		ctx.Set("oh", "yes") //覆盖前面指定KEY
		return ctx.Render(200, "index")
	})

	v.Run(fasthttp.New(":9000"))
}
