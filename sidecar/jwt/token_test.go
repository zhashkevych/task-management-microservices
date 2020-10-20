package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestGenerateAndParseToken(t *testing.T) {
	var accessToken string
	userId := 1

	t.Run("generate token", func(t *testing.T) {
		token := New(TokenInput{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			UserId:    userId,
		})

		tok := jwt.NewWithClaims(jwt.SigningMethodHS256, &token)

		var err error
		accessToken, err = tok.SignedString([]byte(`secret`))
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("parse token", func(t *testing.T) {
		token, err := ParseToken(accessToken)
		if err != nil {
			t.Error(err)
		}

		if token.UserId != userId {
			t.Errorf("user id mismatch")
		}
	})
}
