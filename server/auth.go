package server

import (
	"errors"
	"net/http"

	"github.com/jrandyl/portfolio/token"
	"github.com/labstack/echo/v4"
)

var (
	authorizationPayloadKey = "authorization_payload"
	cookieName              = "user-access"
)

func cookieMiddleware(tokenGenerator token.Generator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cookie, err := c.Cookie(cookieName)

			if err != nil {
				err := errors.New("Unauthorized")
				return c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			}

			acessToken := cookie.Value

			payload, err := tokenGenerator.VerifyToken(acessToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			}

			c.Set(authorizationPayloadKey, payload)
			return next(c)
		}
	}
}
