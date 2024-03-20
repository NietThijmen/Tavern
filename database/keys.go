package database

import (
	"time"
)

var keyCache = make(map[string]bool)

// GetKey retrieves a key from the database and caches it for 5 minutes (used for uploads to the server)
func GetKey(key string) bool {
	if keyCache[key] {
		return keyCache[key]
	}

	row := db.QueryRow("SELECT 1 FROM tavern_tokens WHERE vault_id = ? AND token = ?", vaultID, key)

	var k string
	err := row.Scan(&k)

	if err != nil {
		return false
	}

	keyCache[key] = k != ""

	go func() {
		time.Sleep(5 * time.Minute)
		delete(keyCache, key)
	}()

	return k != ""
}
