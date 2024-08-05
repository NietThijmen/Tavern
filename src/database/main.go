package database

import (
	"database/sql"
	"time"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nietthijmen/tavern/src/config"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
)

var Connection *sql.DB
var Cache = cache.New(5*time.Minute, 10*time.Minute)

func Init() *sql.DB {
	driver := config.ReadEnv("DATABASE_DRIVER", "mysql")
	dsn := config.ReadEnv("DATABASE_DSN", "root@localhost")

	switch driver {
	case "mysql":
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to the database")
		}

		err = db.Ping()
		if err != nil {
			log.Fatal().Err(err).Msg("Error pinging the database")
		}

		Connection = db
	case "sqlite":
		db, err := sql.Open("sqlite", dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Error connecting to the database")
		}

		err = db.Ping()
		if err != nil {
			log.Fatal().Err(err).Msg("Error pinging the database")
		}

		Connection = db
	default:
		log.Fatal().Msg("Database driver is not supported")
	}

	return Connection
}
