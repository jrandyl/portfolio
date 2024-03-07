package server

import (
	"github.com/jrandyl/portfolio/web"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) SetupRouter() {

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.Secure())

	web.RegisterHandler(s.router)

	//pages
	s.router.GET("/", s.index)
	s.router.GET("/projects", s.project)
	s.router.GET("/experience", s.experience)
	s.router.GET("/get-in-touch", s.getInTouch)

	//components
	s.router.GET("/side-nav", s.sideNav)
	s.router.GET("/close-nav", s.closeNav)
}
