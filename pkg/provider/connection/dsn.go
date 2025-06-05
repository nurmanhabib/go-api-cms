package connection

import (
	"fmt"
	"strings"

	"go-api-cms/config"
)

func NewDSN(conf *config.Config) DSN {
	return DSN{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		User:     conf.DB.Username,
		Password: conf.DB.Password,
		DBName:   conf.DB.Database,
		SSLMode:  false,
		Timezone: conf.DB.Timezone,
	}
}

// DSN is a struct for Postgres database connection DSN configuration.
type DSN struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  bool
	Timezone string
}

// Postgres is postgres DSN string format.
func (dsn DSN) Postgres() string {
	var s []string

	if dsn.Host != "" {
		s = append(s, fmt.Sprintf("host=%s", dsn.Host))
	} else {
		s = append(s, fmt.Sprintf("host=%s", "localhost"))
	}

	if dsn.Port > 0 {
		s = append(s, fmt.Sprintf("port=%d", dsn.Port))
	} else {
		s = append(s, fmt.Sprintf("port=%d", 5432))
	}

	if dsn.User != "" {
		s = append(s, fmt.Sprintf("user=%s", dsn.User))
	}

	if dsn.Password != "" {
		s = append(s, fmt.Sprintf("password=%s", dsn.Password))
	}

	if dsn.DBName != "" {
		s = append(s, fmt.Sprintf("dbname=%s", dsn.DBName))
	}

	if dsn.SSLMode {
		s = append(s, fmt.Sprintf("sslmode=%s", "require"))
	} else {
		s = append(s, fmt.Sprintf("sslmode=%s", "disable"))
	}

	if dsn.Timezone != "" {
		s = append(s, fmt.Sprintf("TimeZone=%s", dsn.Timezone))
	}

	return strings.Join(s, " ")
}

// MySQL is MySQL DSN formatted
func (dsn DSN) MySQL() string {
	format := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dsn.User,
		dsn.Password,
		dsn.Host,
		dsn.Port,
		dsn.DBName,
	)

	return format
}
