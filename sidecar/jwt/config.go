package jwt

import "time"

type Config struct {
	Audience string
	Issuer   string
	TTL      time.Duration
}

var cfg Config

func SetConfig(c Config) {
	cfg = c
}
