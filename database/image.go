package database

import "log"

func LogImage(url string) {
	_, err := db.Exec("INSERT INTO tavern_images (vault_id, path) VALUES (?, ?)", vaultID, url)
	if err != nil {
		log.Printf("Error logging image: %s", err)
		return
	} else {
		log.Printf("Logged image: %s", url)
	}
}
