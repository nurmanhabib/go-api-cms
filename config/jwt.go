package config

import (
	"time"

	"go-api-cms/pkg/env"
)

type Auth struct {
	JWTSecret string
	TTL       time.Duration
}

func WithJWT() Option {
	return func(conf *Config) {
		conf.Auth = Auth{
			JWTSecret: env.Get("JWT_SECRET"),
			TTL:       env.GetDuration("JWT_TTL", 15*time.Minute),
		}
	}
}
