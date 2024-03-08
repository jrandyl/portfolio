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

func (s *Server) sccSchool(c echo.Context) error {
	return s.Render(c, components.SCCSchool())
}

func (s *Server) svpdcSchool(c echo.Context) error {
	return s.Render(c, components.SvpdcSchool())
}

func (s *Server) claretSchool(c echo.Context) error {
	return s.Render(c, components.ClaretSchool())
}

func (s *Server) wela(c echo.Context) error {
	return s.Render(c, components.Wela())
}
