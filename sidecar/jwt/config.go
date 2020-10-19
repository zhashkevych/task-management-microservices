package jwt

type Config struct {
	Audience string
	Issuer   string
	Encryption
}

type Encryption struct {
	Algorithm string `json:"alg"`
	Key       string `json:"k"`
}

var cfg Config

func SetConfig(c Config) {
	cfg = c
}