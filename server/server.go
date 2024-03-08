package server

import (
	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/jrandyl/portfolio/database/sqlc"
	"github.com/jrandyl/portfolio/token"
	"github.com/jrandyl/portfolio/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config         utils.Config
	router         *echo.Echo
	store          sqlc.Store
	tokenGenerator token.Generator
}

func NewServer(config utils.Config, transaction sqlc.Store) (*Server, error) {
	tokenGenerator, err := token.NewPasetoGenerator(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		store:          transaction,
		router:         echo.New(),
		tokenGenerator: tokenGenerator,
		config:         config,
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
