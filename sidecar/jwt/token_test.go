package jwt

import (
	"testing"
)

type tokenTest struct {
	name           string
	token          string
	secret         string
	expectedUserId int
}

var (
	testTable = []tokenTest{
		{"correct token",
			"eyJhbGciOiJIUzI1NiIsImtpZCI6InNpbTIifQ.eyJDbGFpbXMiOm51bGwsImF1ZCI6Imh0dHA6Ly9nYXRld2F5OjgwODAiLCJleHAiOjE2MDMxNDIyNjUsImlzcyI6Imh0dHA6Ly91c2Vycy1zZXJ2aWNlOjgwMDAiLCJ1c2VyX2lkIjoxfQ.7GDFBlScMz6BKcZwRMYhy76WKpjTcZVOUliWKVf7bWM",
			"secret",
			1},
	}
)

func TestParseClaims(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Key = tt.secret

			claims, err := ParseClaims(tt.token)
			if err != nil {
				t.Error(err)
			}

			if claims.UserId != tt.expectedUserId {
				t.Errorf("user id mismatch")
			}
		})
	}
}
