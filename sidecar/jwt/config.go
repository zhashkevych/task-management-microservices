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

// SetConfig creates private package-level object of Config type
func SetConfig(c Config) {
	cfg = c
}
