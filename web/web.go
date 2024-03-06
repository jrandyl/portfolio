package web

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:public
	pub embed.FS
	//go:embed public/index.html
	indexHTML embed.FS

	pubDirFS     = echo.MustSubFS(pub, "public")
	pubIndexHTML = echo.MustSubFS(indexHTML, "public/index.html")
)

func RegisterHandler(e *echo.Echo) {
	e.FileFS("/index.html", "index.html", pubIndexHTML)
	e.StaticFS("/", pubDirFS)
}
