package jwt

type Config struct {
	Audience string
	Issuer   string
}

var cfg Config

// SetConfig creates private package-level object of Config type
func SetConfig(c Config) {
	cfg = c
}
