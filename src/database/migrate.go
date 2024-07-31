package database

import "github.com/rs/zerolog/log"

func Migrate() {
	var migrations = []string{
		`CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`,

		`CREATE TABLE IF NOT EXISTS api_keys (
   		id INT AUTO_INCREMENT PRIMARY KEY,
   		user_id INT NOT NULL,
   		api_key VARCHAR(255) NOT NULL,
   		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   		FOREIGN KEY (user_id) REFERENCES users(id)
)`,

		`CREATE TABLE IF NOT EXISTS storage_buckets (
       		id INT AUTO_INCREMENT PRIMARY KEY,
       		name VARCHAR(255) NOT NULL,
       		type VARCHAR(255) NOT NULL,
       		root_path VARCHAR(255) NOT NULL,
       		ip VARCHAR(255) DEFAULT NULL,
       		port INT DEFAULT NULL,
       		username VARCHAR(255) DEFAULT NULL,
       		password VARCHAR(255) DEFAULT NULL,
       		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     )`,

		`CREATE TABLE IF NOT EXISTS storage_objects (
       		id INT AUTO_INCREMENT PRIMARY KEY,
       		bucket_id INT NOT NULL,
       		name VARCHAR(255) NOT NULL,
       		size INT NOT NULL,
       		file_type VARCHAR(255) NOT NULL,
       		file_path VARCHAR(255) NOT NULL,
       		uploaded_by varchar(255) NOT NULL COMMENT 'IP address of the uploader',
       		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       		FOREIGN KEY (bucket_id) REFERENCES storage_buckets(id)
	 )`,
	}

	for _, migration := range migrations {
		_, err := Connection.Exec(migration)
		if err != nil {
			log.Error().Err(err).Msg("Error running migration")
		}

		log.Info().Msg("Migration ran successfully")
	}

	log.Info().Msg("Migrations ran successfully")
}
