package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type (
	AccessToken struct {
		jwt.StandardClaims
		UserId int    `json:"user_id"`
	}

	TokenInput struct {
		UserId    int
		ExpiresAt int64
	}
)

// New generates access token
// Using shared secret from JWK file and config data
// IMPORTANT: SetConfig() and SetEncriptionKeyFromJWK() calls are required before executing this method
func New(input TokenInput) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: input.ExpiresAt,
			Audience:  cfg.Audience,
			Issuer:    cfg.Issuer,
		},
		UserId: input.UserId,
	})

	token.Header["kid"] = cfg.Encryption.KeyId

	return token.SignedString([]byte(cfg.Encryption.Key))
}

// ParseToken extracts payload from access token
// Previously generated with New()
func ParseToken(token string) (AccessToken, error) {
	t, err := jwt.ParseWithClaims(token, &AccessToken{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.Encryption.Key), nil
	})

	if err != nil {
		return AccessToken{}, err
	}

	parsed, ok := t.Claims.(*AccessToken)
	if !ok {
		return AccessToken{}, errors.New("token is of invalid type")
	}

	return *parsed, nil
}
