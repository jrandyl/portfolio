package server

import (
	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

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

func (s *Server) Render(c echo.Context, templ templ.Component) error {
	return templ.Render(c.Request().Context(), c.Response())
}

func (s *Server) Validate(i interface{}) error {
	err := validator.New().Struct(i)
	return err
}

type CustomError struct {
	Error string
}

func ErrorResponse(err error) *CustomError {
	return &CustomError{
		Error: err.Error(),
	}
}
