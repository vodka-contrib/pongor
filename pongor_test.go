package pongor

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/macaron-contrib/pongor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderHtml(t *testing.T) {
	Convey("Render HTML", t, func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		r := GetRenderer(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/echo", func(ctx *echo.Context) error {
			return ctx.Render(http.StatusOK, "echo.html", nil)
		})
		resp := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/echo", nil)
		So(err, ShouldBeNil)
		e.ServeHTTP(resp, req)
		So(resp.Body.String(), ShouldEqual, "<h1>Hello world</h1>\n")
		So(resp.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Render HTML with Context", t, func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		r := GetRenderer(PongorOption{
			Directory: "test",
		})
		e.SetRenderer(r)
		e.Get("/echo", func(ctx *echo.Context) error {
			return ctx.Render(http.StatusOK, "echo_markup.html", map[string]interface{}{
				"name": "echo",
			})
		})
		resp := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/echo", nil)
		So(err, ShouldBeNil)
		e.ServeHTTP(resp, req)
		So(resp.Body.String(), ShouldEqual, "<h1>Hello, echo</h1>\n")
		So(resp.Code, ShouldEqual, http.StatusOK)
	})
}

func ExampleRender() {
	r := pongor.GetRenderer()
	e.SetRenderer(r)
	e.Get("/", func(ctx *echo.Context) error {
		// render ./templates/index.html file.
		ctx.Render(200, "index.html", map[string]interface{}{
			"title": "你好，世界",
		})
		return nil
	})
}
