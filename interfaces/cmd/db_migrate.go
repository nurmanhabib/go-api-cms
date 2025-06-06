package cmd

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v3"

	"go-api-cms/internal/basepath"
)

func DBMigrate(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:  "db:migrate",
		Usage: "Run database migrations",
		Action: func(ctx context.Context, command *cli.Command) error {
			if err := goose.SetDialect("postgres"); err != nil {
				return err
			}

			if err := goose.Up(db, basepath.Dir("database/migrations")); err != nil {
				return err
			}

			return nil
		},
	}
}

func DBRollback(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:  "db:rollback",
		Usage: "Run database migrations",
		Action: func(ctx context.Context, command *cli.Command) error {
			if err := goose.SetDialect("postgres"); err != nil {
				return err
			}

			if err := goose.Down(db, basepath.Dir("database/migrations")); err != nil {
				return err
			}

			return nil
		},
	}
}
