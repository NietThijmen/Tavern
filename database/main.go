package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nietthijmen/tavern/config"
	"log"
)

var db *sql.DB

// Init connects to the database and sets the global db variable
func Init() {
	var username = config.ReadEnv("DB_USERNAME", "root")
	var password = config.ReadEnv("DB_PASSWORD", "root")
	var host = config.ReadEnv("DB_HOST", "localhost")
	var port = config.ReadEnv("DB_PORT", "3306")
	var database = config.ReadEnv("DB_DATABASE", "tavern")

	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 host + ":" + port,
		DBName:               database,
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to database %s on %s:%s", database, host, port)
}
