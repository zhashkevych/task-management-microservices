package jwt

import (
	"encoding/json"
	"testing"
	"time"
)

// TODO refactor tests

type tokenTest struct {
	name           string
	token          string
	secret         string
	expectedUserId int
	shouldFail bool
}

var (
	testTable = []tokenTest{
		{"correct token",
			"eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ.eyJhdWQiOiJodHRwOi8vZ2F0ZXdheTo4MDgwIiwiZXhwIjoxNjAzMTQyMjY1LCJpc3MiOiJodHRwOi8vdXNlcnMtc2VydmljZTo4MDAwIiwidXNlcl9pZCI6MX0.s1goU57bVF3ViPAPncsEtR1MogREtfHHIdOa28jKHoY",
			"secret",
			1,
		false},
		{"incorrect token",
			"eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ.eyJhdWQiOiJodHRwOi8vZ2F0ZXdheTo4MDgwIiwiZXhwIjoxNjAzMTQyMjY1LCJpc3MiOiJodHRwOi8vdXNlcnMtc2VydmljZTo4MDAwIiwidXNlcl9pZCI6MX0.s1goU57bVF3ViPAPncsEtR1MogREtfHHIdOa28jY",
			"secret",
			1,
			true},
	}
)

func TestNewAccessToken(t *testing.T) {
	t.Run("generate token", func(t *testing.T) {
		token := NewAccessToken(TokenInput{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			UserId: 1,
		})

		str, _ := json.Marshal(token)
		t.Log("token: ", string(str))
	})
}

func TestParseToken(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Encryption.Key = tt.secret

			token, err := ParseToken(tt.token)
			if err != nil && tt.shouldFail {
				t.Skip("ok")
			}

			if token.UserId != tt.expectedUserId {
				t.Errorf("user id mismatch")
			}
		})
	}
}
