package server

import (
	"github.com/jrandyl/portfolio/client"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) SetupRouter() {

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.Secure())

	client.RegisterHandler(s.router)

	// s.router.GET("/projects", s.project)
	// s.router.GET("/about", s.about)
	// s.router.GET("/product", s.product)

	// s.router.POST("/login", s.userLogin)
	// app := s.router.Group("/app")
	// app.GET("/dashboard", s.index)
}
