package jwt

import (
	"errors"
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

// New generates AccessToken instance
// Which is used as payload for JWT, signed on API Gateway
func New(input TokenInput) AccessToken {
	return AccessToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: input.ExpiresAt,
			Audience:  cfg.Audience,
			Issuer:    cfg.Issuer,
		},
		UserId: input.UserId,
	}
}

// ParseToken extracts payload from access token
// Previously generated with New()
func ParseToken(token string) (AccessToken, error) {
	t, _, err := new(jwt.Parser).ParseUnverified(token, &AccessToken{})
	if err != nil {
		return AccessToken{}, err
	}

	claims, ok := t.Claims.(*AccessToken)
	if !ok {
		return AccessToken{}, errors.New("incorrect claims type")
	}

	return *claims, nil
}
