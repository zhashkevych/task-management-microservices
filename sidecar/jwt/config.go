package jwt

type Config struct {
	Audience string
	Issuer   string
	Encryption
}

type Encryption struct {
	Algorithm string `json:"alg"`
	Key       string `json:"k"`
	KeyId     string `json:"kid"`
}

var cfg Config

func SetConfig(c Config) {
	cfg = c
}
