package token

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Generator interface {
	CreateToken(userID string, duration time.Duration) (string, error)
	CreateTokenAndCookie(userID string, duration time.Duration, c echo.Context) error
	VerifyToken(token string) (*Payload, error)
}
