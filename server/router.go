package server

import (
	"github.com/jrandyl/portfolio/web"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) SetupRouter() {
	auth := cookieMiddleware(s.tokenGenerator)

	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.Secure())

	web.RegisterHandler(s.router)

	//pages
	s.router.GET("/", s.index)
	s.router.GET("/svpdc-school", s.svpdcSchool)
	s.router.GET("/wela", s.wela)
	s.router.GET("/claret-school", s.claretSchool)
	s.router.GET("/scc-school", s.sccSchool)
	s.router.GET("/projects", s.project)
	s.router.GET("/experience", s.experience)
	s.router.GET("/get-in-touch", s.getInTouch)
	s.router.POST("/message/client", s.client)

	//secret pages
	secret := s.router.Group("secret")
	secret.GET("/register", s.register)
	secret.POST("/user/register", s.userRegister)
	secret.GET("/login", s.login)
	secret.POST("/login", s.userLogin)

	//app
	app := s.router.Group("app")
	app.Use(auth)
	app.GET("/dashboard", s.dashboard)

	//components
	s.router.GET("/side-nav", s.sideNav)
	s.router.GET("/close-nav", s.closeNav)
}
