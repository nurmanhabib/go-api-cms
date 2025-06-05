package config

import "go-api-cms/pkg/env"

type Database struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Timezone string
}

func WithDatabase() Option {
	return func(conf *Config) {
		conf.DB = Database{
			Driver:   env.Get("DB_DRIVER", "postgres"),
			Host:     env.Get("DB_HOST", "127.0.0.1"),
			Port:     env.GetInt("DB_PORT", 5432),
			Username: env.Get("DB_USERNAME", "root"),
			Password: env.Get("DB_PASSWORD", ""),
			Database: env.Get("DB_DATABASE", "go_api"),
			Timezone: env.Get("DB_TIMEZONE", "Asia/Jakarta"),
		}
	}
}
