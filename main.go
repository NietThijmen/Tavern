package main

import (
	"github.com/nietthijmen/tavern/src/database"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	database.Init()
	app := &cli.App{
		Name:  "tavern",
		Usage: "Finally, an open source scalable storage solution",
		Commands: []*cli.Command{
			// Database migration command
			{
				Name:     "database:migrate",
				Usage:    "Migrate the database",
				Category: "Database",
				Action: func(c *cli.Context) error {
					database.Migrate()
					return nil
				},
			},
			// Database status command
			{
				Name:     "database:status",
				Usage:    "Check the database status",
				Category: "Database",
				Action: func(c *cli.Context) error {
					if database.Connection.Ping() == nil {
						log.Info().Msg("Database connection is successful")
					} else {
						log.Error().Msg("Database connection failed")
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Error running the application")
	}
}
