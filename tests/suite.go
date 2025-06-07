package tests

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"

	"go-api-cms/app"
	"go-api-cms/config"
	"go-api-cms/database/seeder"
	"go-api-cms/interfaces/cmd"
	"go-api-cms/internal/basepath"
	"go-api-cms/tests/database"
)

type Suite struct {
	App *app.App
	Ctx context.Context
}

func NewSuite() *Suite {
	_ = godotenv.Load(basepath.Dir(".env"))
	ctx := context.Background()

	conf := config.New(
		config.WithTestDatabase(),
		config.WithJWT(),
	)

	apps := app.New(
		app.WithConfig(conf),
		app.WithDatabase(),
		app.WithRepositories(),
	)

	commands := &cli.Command{
		Commands: []*cli.Command{
			cmd.DBMigrate(apps.Sql),
			cmd.DBSeeder(apps.DB),
		},
	}

	// Prepare clear database first
	switch apps.Config.DB.Driver {
	case "postgres":
		if errDrop := database.ClearPostgresSchema(apps.Sql); errDrop != nil {
			panic(errDrop)
		}

	case "mysql":
		if errDrop := database.ClearMySQLDatabase(apps.Sql, apps.Config.DB.Database); errDrop != nil {
			panic(errDrop)
		}
	}

	if errRun := commands.Run(ctx, []string{"", "db:migrate"}); errRun != nil {
		panic(errRun)
	}

	if errRun := commands.Run(ctx, []string{"", "db:seed"}); errRun != nil {
		panic(errRun)
	}

	return &Suite{
		App: apps,
		Ctx: ctx,
	}
}

func (s *Suite) RunSeeder(seeders ...seeder.Seeder) error {
	for _, seed := range seeders {
		if err := seed.Seed(s.Ctx, s.App.DB); err != nil {
			return err
		}
	}
	return nil
}
