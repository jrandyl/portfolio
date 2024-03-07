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

func (s *Server) index(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Index())
	}

	return s.Render(c, templates.Index())
}

func (s *Server) experience(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Experience())
	}

	return s.Render(c, templates.Experience())
}

func (s *Server) getInTouch(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.GetInTouch())
	}

	return s.Render(c, templates.GetInTouch())
}
