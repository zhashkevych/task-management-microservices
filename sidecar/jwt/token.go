package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type (
	Claims struct {
		jwt.StandardClaims
		UserId int `json:"user_id"`
	}

	AccessToken struct {
		Audience  string `json:"aud"`
		Issuer    string `json:"iss"`
		Claims
	}
)

func NewAccessToken(claims Claims) AccessToken {
	return AccessToken{
		Audience:  cfg.Audience,
		Issuer:    cfg.Issuer,
		Claims:    claims,
	}
}

func ParseClaims(token string) (Claims, error) {
	t, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		//TODO check signing method by alg provided in JWK
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return cfg.Encryption.Key, nil
	})

	if err != nil {
		return Claims{}, err
	}

	claims, ok := t.Claims.(Claims)
	if !ok {
		return Claims{}, errors.New("failed to get claims from token")
	}

	return claims, nil
}
