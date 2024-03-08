package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jrandyl/portfolio/database/sqlc"
	"github.com/jrandyl/portfolio/web/src/components"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (s *Server) client(c echo.Context) error {
	first_name := c.FormValue("first_name")
	middle_name := c.FormValue("middle_name")
	last_name := c.FormValue("last_name")
	email := c.FormValue("email")
	message := c.FormValue("message")

	id := uuid.New().String()

	arg := sqlc.CreateClientParams{
		ID:         id,
		FirstName:  first_name,
		MiddleName: middle_name,
		LastName:   last_name,
		FullName:   first_name + " " + middle_name + " " + last_name,
		Email:      email,
		Message:    message,
	}

	_, err := s.store.CreateClient(c.Request().Context(), arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.JSON(http.StatusForbidden, ErrorResponse(err))
			}
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Unable to register new account: %v", ErrorResponse(err)))
	}

	return s.Render(c, components.Success())
}
