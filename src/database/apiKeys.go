package database

import "github.com/google/uuid"

func GetApiKey(apiKey string) (ApiKey, error) {
	var key ApiKey

	err := Connection.QueryRow("SELECT * FROM api_keys WHERE api_key = ?", apiKey).Scan(&key.Id, &key.ApiKey, &key.CreatedAt)
	if err != nil {
		return key, err
	}

	return key, nil
}
func GetApiKeys() ([]ApiKey, error) {
	var keys []ApiKey

	rows, err := Connection.Query("SELECT * FROM api_keys")
	if err != nil {
		return keys, err
	}

	for rows.Next() {
		var key ApiKey
		err := rows.Scan(&key.Id, &key.ApiKey, &key.CreatedAt)
		if err != nil {
			return keys, err
		}

		keys = append(keys, key)
	}

	return keys, nil
}
func GenerateApiKey() (ApiKey, error) {
	newID := uuid.New()

	_, err := Connection.Exec("INSERT INTO api_keys (api_key) VALUES (?)", "tavern_"+newID.String())
	if err != nil {
		return ApiKey{}, err
	}

	return GetApiKey("tavern_" + newID.String())
}
func DeleteApiKey(apiKey string) error {
	_, err := Connection.Exec("DELETE FROM api_keys WHERE api_key = ?", apiKey)
	if err != nil {
		return err
	}

	return nil
}
