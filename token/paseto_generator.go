package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/labstack/echo/v4"
	"github.com/o1egl/paseto"
)

type PasetoGenerator struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoGenerator(symmetricKey string) (Generator, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	generator := &PasetoGenerator{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return generator, nil
}

func (generator *PasetoGenerator) CreateToken(userID string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	return generator.paseto.Encrypt(generator.symmetricKey, payload, nil)
}

func (generator *PasetoGenerator) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := generator.paseto.Decrypt(token, generator.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (generator *PasetoGenerator) CreateTokenAndCookie(userID string, duration time.Duration, c echo.Context) error {
	token, err := generator.CreateToken(userID, duration)
	if err != nil {
		return err
	}

	exp := time.Now().Add(duration)

	generateCookie("user-access", token, exp, c)

	return nil
}
