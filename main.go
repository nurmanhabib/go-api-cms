package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v3"

	"go-api-cms/interfaces/cmd"
	"go-api-cms/pkg/graceful"
	"go-api-cms/routes"
)

func main() {
	commands := &cli.Command{
		Commands: []*cli.Command{
			cmd.DBMigrate(),
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
