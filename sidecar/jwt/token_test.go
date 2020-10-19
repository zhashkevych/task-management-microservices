package jwt

import (
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
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDMxNDk0OTgsInVzZXJfaWQiOjF9.IfxQXrIZFeeQ9gHtkaRl9qoRJZhrMm3wtFDXvZdsu6s",
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

func TestNew(t *testing.T) {
	t.Run("generate token", func(t *testing.T) {
		cfg.Encryption.Key = "secret"

		token, err := New(TokenInput{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			UserId: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log("token: ", token)
	})
}

func TestParseToken(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Encryption.Key = tt.secret

			token, err := ParseToken(tt.token)
			if err != nil {
				if tt.shouldFail {
					t.Skip("ok")
				}

				t.Error(err)
			}

			if token.UserId != tt.expectedUserId {
				t.Errorf("user id mismatch")
			}
		})
	}
}
