package web

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:public
	pub embed.FS

	pubDirFS = echo.MustSubFS(pub, "public")
)

func RegisterHandler(e *echo.Echo) {
	e.StaticFS("/*", pubDirFS)
}
