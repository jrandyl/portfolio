package server

import (
	"github.com/jrandyl/portfolio/web/src/pages/fragments"
	"github.com/jrandyl/portfolio/web/src/pages/templates"
	"github.com/labstack/echo/v4"
)

func (s *Server) project(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Project())
	}

	return s.Render(c, templates.Project())
}
