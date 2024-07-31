package main

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/nietthijmen/tavern/src/database"
	"github.com/nietthijmen/tavern/src/encryption"
	"github.com/nietthijmen/tavern/src/input"
	"github.com/nietthijmen/tavern/src/storage"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
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
			// Generate encryption key
			{
				Name:     "encryption:generate",
				Usage:    "Generate a new encryption key",
				Category: "Encryption",
				Action: func(c *cli.Context) error {
					key := make([]byte, 32)

					_, err := rand.Read(key)
					if err != nil {
						log.Error().Err(err).Msg("Error generating the encryption key")
					} else {
						log.Info().Msg("Generated encryption key")
						encoded := base64.StdEncoding.EncodeToString(key)
						println(encoded)
					}

					return nil
				},
			},
			// Test encryption
			{
				Name:     "encryption:test",
				Usage:    "Test the encryption",
				Category: "Encryption",
				Action: func(c *cli.Context) error {
					encrypted, err := encryption.Encrypt(c.Args().First())
					if err != nil {
						log.Error().Err(err).Msg("Error encrypting the value")
					} else {
						println("Encrypted value:")
						println(encrypted)
					}

					decrypted, err := encryption.Decrypt(encrypted)
					if err != nil {
						log.Error().Err(err).Msg("Error decrypting the value")
					} else {
						println("Decrypted value:")
						println(decrypted)
					}

					return nil
				},
			},
			// Get bucket drivers
			{
				Name:     "bucket:drivers",
				Usage:    "List all available bucket drivers",
				Category: "Bucket",
				Action: func(c *cli.Context) error {
					for _, driver := range storage.DriverTypes {
						println(driver)
					}

					return nil

				},
			},
			// Create a new bucket
			{
				Name:     "bucket:create",
				Usage:    "Create a new bucket",
				Category: "Bucket",
				Action: func(c *cli.Context) error {
					var toCreate database.StorageBucket
					var bucketType string
					bucketType = input.Select(storage.DriverTypes)
					if bucketType == "" {
						log.Error().Msg("Invalid driver type")
						return nil
					}

					toCreate.Type = bucketType

					for _, field := range storage.DriverFields[bucketType] {
						toCreateField := input.Question(field + ":")
						if toCreateField == "" {
							log.Error().Msg("Invalid " + field)
							return nil
						}

						switch field {
						case "Name":
							toCreate.Name = toCreateField
						case "MaxSize":
							asInt, err := strconv.Atoi(toCreateField)
							if err != nil {
								log.Error().Msg("Invalid " + field)
								return nil
							}

							toCreate.MaxSize = asInt

						case "RootPath":
							toCreate.RootPath = toCreateField
						case "Ip":
							toCreate.Ip = toCreateField
						case "Port":
							asInt, err := strconv.Atoi(toCreateField)
							if err != nil {
								log.Error().Msg("Invalid " + field)
								return nil
							}

							toCreate.Port = asInt
						case "Username":
							toCreate.Username = toCreateField
						case "Password":
							encrypted, err := encryption.Encrypt(toCreateField)
							if err != nil {
								log.Error().Err(err).Msg("Error encrypting the password")
								return nil
							}

							toCreate.Password = encrypted
						}
					}

					err := database.CreateBucket(toCreate)
					if err != nil {
						log.Error().Err(err).Msg("Error creating the bucket")
					} else {
						log.Info().Msg("Bucket created successfully")
					}

					return nil
				},
			},
			// List all buckets
			{
				Name:     "bucket:list",
				Usage:    "List all buckets",
				Category: "Bucket",
				Action: func(c *cli.Context) error {
					buckets, err := database.GetAllBuckets()
					if err != nil {
						log.Error().Err(err).Msg("Error listing the buckets")
					} else {
						for _, bucket := range buckets {
							println(bucket.Name + " (" + bucket.Type + ")")
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
