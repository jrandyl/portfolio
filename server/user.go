package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/jrandyl/portfolio/database/sqlc"
	"github.com/jrandyl/portfolio/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (s *Server) userRegister(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	first_name := c.FormValue("first_name")
	middle_name := c.FormValue("middle_name")
	last_name := c.FormValue("last_name")
	email := c.FormValue("email")

	hashedPassword, err := utils.Hash(string(password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Unable to hash password: %v", ErrorResponse(err)))
	}

	arg := sqlc.RegisterTxParams{
		Username:      username,
		Password:      hashedPassword,
		FirstName:     first_name,
		MiddleName:    middle_name,
		LastName:      last_name,
		FullName:      first_name + " " + middle_name + " " + last_name,
		Email:         email,
		ContactNumber: "",
	}

	_, err = s.store.RegisterTx(c.Request().Context(), arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.JSON(http.StatusForbidden, ErrorResponse(err))
			}
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Unable to register new account: %v", ErrorResponse(err)))
	}

	return c.Redirect(http.StatusMovedPermanently, "/secret/login")
}

func (s *Server) userLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := s.store.GetUserUsername(c.Request().Context(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "User not found")
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Cannot fetch user: %v", ErrorResponse(err)))
	}

	err = utils.PasswordChecker(password, user.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Wrong password: %v", ErrorResponse(err)))
	}

	err = s.tokenGenerator.CreateTokenAndCookie(
		user.ID,
		s.config.AccessTokenDuration,
		c,
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	arg := sqlc.UpdateUserStatusParams{
		ID:        user.ID,
		Status:    "Online",
		LastLogin: time.Now().Format("2006-01-02 15:04:05.999999"),
	}

	user, err = s.store.UpdateUserStatus(c.Request().Context(), arg)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Cannot authenticate user: %v", ErrorResponse(err)))
	}

	return c.Redirect(http.StatusMovedPermanently, "/app/dashboard")
}
