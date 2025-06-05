package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-api-cms/config"
)

func NewDBConnection(conf *config.Config) (*gorm.DB, error) {
	dsn := NewDSN(conf)
	switch conf.DB.Driver {
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn.MySQL()), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil

	case "postgres":
		db, err := gorm.Open(postgres.Open(dsn.Postgres()), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil

	default:
		return nil, fmt.Errorf("unknown database driver: %s", conf.DB.Driver)
	}
}
