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

			// Generate API key
			{
				Name:     "key:generate",
				Usage:    "Generate a new API key",
				Category: "Api key",
				Action: func(c *cli.Context) error {
					key, err := database.GenerateApiKey()
					if err != nil {
						log.Error().Err(err).Msg("Error generating the API key")
					} else {
						log.Info().Msg("Generated API key: " + key.ApiKey)
					}

					return nil
				},
			},

			// Delete API key
			{
				Name:     "key:delete",
				Usage:    "Delete an API key",
				Category: "Api key",
				Action: func(c *cli.Context) error {
					err := database.DeleteApiKey(c.Args().First())
					if err != nil {
						log.Error().Err(err).Msg("Error deleting the API key")
					} else {
						log.Info().Msg("Deleted API key: " + c.Args().First())
					}

					return nil
				},
			},

			// List API keys
			{
				Name:     "key:list",
				Usage:    "List all API keys",
				Category: "Api key",
				Action: func(c *cli.Context) error {
					keys, err := database.GetApiKeys()
					if err != nil {
						log.Error().Err(err).Msg("Error listing the API keys")
					} else {
						for _, key := range keys {
							println(key.ApiKey)
						}
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
