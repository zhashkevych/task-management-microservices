package jwt

import "time"

type (
	Claims struct {
		UserId int `json:"user_id"`
	}

	AccessToken struct {
		Audience  string `json:"aud"`
		Issuer    string `json:"iss"`
		ExpiresAt int64  `json:"exp"`
		Claims
	}

	RefreshToken struct {
	}
)

func NewAccessToken(claims Claims) AccessToken {
	return AccessToken{
		Audience:  cfg.Audience,
		Issuer:    cfg.Issuer,
		ExpiresAt: time.Now().Add(cfg.TTL).Unix(),
		Claims:    claims,
	}
}
