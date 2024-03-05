package server

import "github.com/labstack/echo/v4"

type Server struct {
	router *echo.Echo
}

func NewServer() (*Server, error) {

	server := &Server{
		router: echo.New(),
	}

	server.SetupRouter()

	return server, nil
}

func (s *Server) Start(port string) error {
	return s.router.Start(port)
}
