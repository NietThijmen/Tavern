package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nietthijmen/tavern/src/config"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"time"
)

var Connection *sql.DB
var Cache = cache.New(5*time.Minute, 10*time.Minute)

func Init() *sql.DB {
	db, err := sql.Open(config.ReadEnv("DATABASE_DRIVER", "mysql"), config.ReadEnv("DATABASE_DSN", "root@localhost"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to the database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Error pinging the database")
	}

	Connection = db

	return Connection
}
