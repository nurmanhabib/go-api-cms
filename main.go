package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"

	"go-api-cms/app"
	"go-api-cms/config"
	"go-api-cms/interfaces/cmd"
	"go-api-cms/pkg/graceful"
	"go-api-cms/routes"
)

func main() {
	_ = godotenv.Load(".env")

	conf := config.New(
		config.WithDatabase(),
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
			cmd.DBRollback(apps.Sql),
			cmd.DBSeeder(apps.DB),
		},

		// Default HTTP Server
		Action: func(ctx context.Context, command *cli.Command) error {
			router := routes.Api(apps)
			return graceful.RunHTTPServerWithGracefulShutdown(router, fmt.Sprintf(":%d", apps.Config.App.Port), 10*time.Second)
		},
	}

	if err := commands.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
