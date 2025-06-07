package app

import (
	"fmt"

	"go-api-cms/config"
	"go-api-cms/infrastructure/dao"
	"go-api-cms/pkg/provider/connection"
)

func WithConfig(conf *config.Config) Option {
	return func(app *App) error {
		app.Config = conf
		return nil
	}
}

func WithDatabase() Option {
	return func(app *App) error {
		if app.Config == nil {
			return fmt.Errorf("config is nil")
		}

		dbConn, errDBConn := connection.NewDBConnection(app.Config)
		if errDBConn != nil {
			return errDBConn
		}

		sqlDB, errSqlDB := dbConn.DB()
		if errSqlDB != nil {
			return errSqlDB
		}

		app.DB = dbConn
		app.Sql = sqlDB

		return nil
	}
}

func WithRepositories() Option {
	return func(app *App) error {
		if app.DB == nil {
			return fmt.Errorf("database connection is nil")
		}

		repo := dao.NewRepositories(app.DB)
		app.Repo = repo
		return nil
	}
}
