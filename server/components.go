package server

import (
	"github.com/jrandyl/portfolio/web/src/components"
	"github.com/labstack/echo/v4"
)

func (s *Server) sideNav(c echo.Context) error {
	return s.Render(c, components.SideNav())
}

func (s *Server) closeNav(c echo.Context) error {
	return s.Render(c, components.CloseNav())
}
