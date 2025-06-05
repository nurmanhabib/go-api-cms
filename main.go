package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"

	"go-api-cms/config"
	"go-api-cms/interfaces/cmd"
	"go-api-cms/pkg/graceful"
	"go-api-cms/pkg/provider/connection"
	"go-api-cms/routes"
)

func main() {
	_ = godotenv.Load(".env")

	conf := config.New(
		config.WithDatabase(),
	)

	dbConn, errDBConn := connection.NewDBConnection(conf)
	if errDBConn != nil {
		panic(errDBConn)
	}

	sqlDB, errSqlDB := dbConn.DB()
	if errSqlDB != nil {
		panic(errSqlDB)
	}

	commands := &cli.Command{
		Commands: []*cli.Command{
			cmd.DBMigrate(sqlDB),
		},

		// Default HTTP Server
		Action: func(ctx context.Context, command *cli.Command) error {
			router := routes.Api()
			return graceful.RunHTTPServerWithGracefulShutdown(router, ":8087", 10*time.Second)
		},
	}

	if err := commands.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
