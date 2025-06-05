package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
)

func DBMigrate() *cli.Command {
	return &cli.Command{
		Name:  "db:migrate",
		Usage: "Run database migrations",
		Action: func(ctx context.Context, command *cli.Command) error {
			return nil
		},
	}
}
