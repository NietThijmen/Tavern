package database

import (
	"github.com/nietthijmen/tavern/config"
	"log"
	"time"
)

var keyCache = make(map[string]string)
var vaultID string = config.ReadEnv("VAULT_ID", "1")

func GetKey(key string) string {
	log.Printf("Getting key: %s", key)
	log.Printf("Vault ID: %s", vaultID)

	if keyCache[key] != "" {
		return keyCache[key]
	}

	row := db.QueryRow("SELECT 1 FROM tavern_tokens WHERE vault_id = ? AND token = ?", vaultID, key)

	var k string
	err := row.Scan(&k)

	if err != nil {
		return ""
	}

	keyCache[key] = k

	go func() {
		time.Sleep(5 * time.Minute)
		delete(keyCache, key)
	}()

	return k
}
