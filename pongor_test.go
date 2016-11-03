package pongor

import (
	"net/http"
	"testing"

	"github.com/insionng/vodka/test"

	"github.com/insionng/vodka"
	. "github.com/smartystreets/goconvey/convey"
)

func request(method, path string, e *vodka.Vodka) (int, string) {
	req := test.NewRequest(method, path, nil)
	rec := test.NewResponseRecorder()
	e.ServeHTTP(req, rec)
	return rec.Status(), rec.Body.String()
}

func TestRenderHtml(t *testing.T) {
	Convey("Render HTML", t, func() {
		e := vodka.New()
		r := Renderor(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(ctx vodka.Context) error {
				return ctx.Render(http.StatusOK, "vodka")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello vodka</h1>")
	})

	Convey("Render HTML with Reload", t, func() {
		e := vodka.New()
		r := Renderor(PongorOption{
			Directory: "test",
			Reload:    true,
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(ctx vodka.Context) error {
				return ctx.Render(http.StatusOK, "vodka")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello vodka</h1>")
	})

	Convey("Render HTML with Context", t, func() {
		e := vodka.New()
		r := Renderor(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(ctx vodka.Context) error {
				ctx.Set("name", "vodka")
				return ctx.Render(http.StatusOK, "vodka_markup")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, vodka</h1>")
	})

	Convey("Render HTML with Context and Reload", t, func() {
		e := vodka.New()
		r := Renderor(PongorOption{
			Directory: "test",
			Reload:    true,
		})
		e.SetRenderer(r)
		e.Get("/vodka", func() vodka.HandlerFunc {
			return func(ctx vodka.Context) error {
				ctx.Set("name", "vodka")
				return ctx.Render(http.StatusOK, "vodka_markup")
			}
		}())
		status, body := request("GET", "/vodka", e)
		So(status, ShouldEqual, http.StatusOK)
		So(body, ShouldEqual, "<h1>Hello, vodka</h1>")
	})
}

func ExampleRender() {
	e := vodka.New()
	r := Renderor()
	e.SetRenderer(r)
	e.Get("/", func() vodka.HandlerFunc {
		return func(ctx vodka.Context) error {
			ctx.Set("title", "你好，世界")
			// render ./templates/index file.
			ctx.Render(200, "index")
			return nil
		}
	}())
}
