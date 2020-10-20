package jwt

import (
	"testing"
	"time"
)

func TestGenerateAndParseToken(t *testing.T) {
	var accessToken string
	userId := 1
	secret := "secret_key"

	t.Run("generate token", func(t *testing.T) {
		cfg.Encryption.Key = secret

		token, err := New(TokenInput{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			UserId: userId,
		})
		if err != nil {
			t.Fatal(err)
		}

		accessToken = token
	})

	t.Run("parse token", func(t *testing.T) {
		cfg.Encryption.Key = secret

		token, err := ParseToken(accessToken)
		if err != nil {
			t.Error(err)
		}

		if token.UserId != userId {
			t.Errorf("user id mismatch")
		}
	})
}
