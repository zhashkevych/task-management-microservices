package domain

type (
	AccessToken struct {
		Audience  string `json:"aud"`
		Issuer    string `json:"iss"`
		Subject   string `json:"sub"`
		ExpiresAt int64  `json:"exp"`
	}

	RefreshToken struct {
	}
)
