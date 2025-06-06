package config

import "go-api-cms/pkg/env"

type Config struct {
	App  App
	DB   Database
	Auth Auth
}

type Option func(conf *Config)

func New(options ...Option) *Config {
	conf := &Config{
		App: App{
			Name:     env.Get("APP_NAME", "go-api-cms"),
			Port:     env.GetInt("APP_PORT", 8080),
			Timezone: env.Get("APP_TIMEZONE", "Asia/Jakarta"),
		},
	}

	for _, option := range options {
		option(conf)
	}

	return conf
}
