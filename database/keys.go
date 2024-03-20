package database

// GetKey retrieves a key from the database and returns a bool if it exists
func GetKey(key string) bool {
	row := db.QueryRow("SELECT 1 FROM tavern_tokens WHERE vault_id = ? AND token = ?", vaultID, key)

	var k string
	err := row.Scan(&k)

	if err != nil {
		return false
	}

	return k != ""
}
