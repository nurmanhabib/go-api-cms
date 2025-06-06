package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
	"gorm.io/gorm"

	"go-api-cms/database/seeder"
)

func DBSeeder(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name: "db:seed",
		Action: func(ctx context.Context, command *cli.Command) error {
			s := seeder.NewRegistry(
				seeder.NewPermissionGroupSeeder(),
				seeder.NewPermissionSeeder(),
				seeder.NewRoleSeeder(),
			)

			return s.Run(ctx, db)
		},
	}
}
