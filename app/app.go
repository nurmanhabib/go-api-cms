package app

import (
	"database/sql"

	"gorm.io/gorm"

	"go-api-cms/config"
	"go-api-cms/infrastructure/dao"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Sql    *sql.DB
	Repo   *dao.Repositories
}

type Option func(app *App) error

func New(options ...Option) *App {
	app := &App{}
	for _, option := range options {
		err := option(app)
		if err != nil {
			return nil
		}
	}
	return app
}
