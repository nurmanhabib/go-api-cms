package config

import (
	"fmt"

	"go-api-cms/pkg/env"
)

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
			Database: env.Get("DB_DATABASE", "go_api_cms"),
			Timezone: env.Get("DB_TIMEZONE", "Asia/Jakarta"),
		}
	}
}

func WithTestDatabase() Option {
	return func(conf *Config) {
		conf.DB = Database{
			Driver:   env.Get("TEST_DB_DRIVER", env.Get("DB_DRIVER", "postgres")),
			Host:     env.Get("TEST_DB_HOST", env.Get("DB_HOST", "127.0.0.1")),
			Port:     env.GetInt("TEST_DB_PORT", env.GetInt("DB_PORT", 5432)),
			Username: env.Get("TEST_DB_USERNAME", env.Get("DB_USERNAME", "rootSS")),
			Password: env.Get("TEST_DB_PASSWORD", env.Get("DB_PASSWORD", "")),
			Database: env.Get("TEST_DB_DATABASE", fmt.Sprintf("%s_test", env.Get("DB_DATABASE", "go_api_cms"))),
			Timezone: env.Get("TEST_DB_TIMEZONE", "Asia/Jakarta"),
		}
	}
}
