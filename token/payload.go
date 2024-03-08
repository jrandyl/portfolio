package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidToken error = errors.New("invalid token")
	ErrExpiredToken error = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	CompanyID string    `json:"company_id"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

func generateCookie(name, token string, exp time.Time, c echo.Context) {
	cookie := new(http.Cookie)

	cookie.Name = name
	cookie.Value = token
	cookie.Expires = exp
	cookie.HttpOnly = true
	cookie.Path = "/"

	c.SetCookie(cookie)
}
