package server

import (
	"fmt"
	"net/http"

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

func (s *Server) login(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Login())
	}

	return s.Render(c, templates.Login())
}

func (s *Server) dashboard(c echo.Context) error {
	clients, err := s.store.ListClients(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Cannot get the list of clients: %v", ErrorResponse(err)))
	}

	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Dashboard(clients))
	}

	return s.Render(c, templates.Dashboard(clients))
}

func (s *Server) register(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.Register())
	}

	return s.Render(c, templates.Register())
}

func (s *Server) getInTouch(c echo.Context) error {
	isRequestHX := c.Request().Header.Get("Hx-Request") != ""

	if isRequestHX {
		return s.Render(c, fragments.GetInTouch())
	}

	return s.Render(c, templates.GetInTouch())
}
