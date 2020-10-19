package jwt

import (
	"github.com/zhashkevych/task-management-microservices/users-service/internal/config"
	"github.com/zhashkevych/task-management-microservices/users-service/internal/domain"
	"time"
)

type Issuer struct {
	audience string
	issuer   string
	ttl      time.Duration
}

func NewIssuer(cfg config.Config) *Issuer {
	return &Issuer{
		audience: cfg.Token.Audience,
		issuer:   cfg.Token.Issuer,
		ttl:      cfg.Token.TTL,
	}
}

func (i *Issuer) Issue(subject string) domain.AccessToken {
	return domain.AccessToken{
		Audience:  i.audience,
		Issuer:    i.issuer,
		Subject:   subject,
		ExpiresAt: time.Now().Add(i.ttl).Unix(),
	}
}
