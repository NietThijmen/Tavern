package database

import (
	"errors"
	"github.com/nietthijmen/tavern/src/encryption"
	"github.com/rs/zerolog/log"
	"math/rand"
)

func GetAllBuckets() ([]StorageBucket, error) {
	query, err := Connection.Query("SELECT * FROM storage_buckets")
	if err != nil {
		return nil, err
	}

	var buckets []StorageBucket
	for query.Next() {
		var bucket StorageBucket
		err = query.Scan(&bucket.Id, &bucket.Name, &bucket.Type, &bucket.MaxSize, &bucket.RootPath, &bucket.Ip, &bucket.Port, &bucket.Username, &bucket.Password, &bucket.CreatedAt)
		if err != nil {
			return nil, err
		}

		if bucket.Password != "" {
			decrypted, err := encryption.Decrypt(bucket.Password)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to decrypt password")
			}

			bucket.Password = decrypted
		}

		buckets = append(buckets, bucket)
	}

	return buckets, nil
}
func GetBucketById(id int) (StorageBucket, error) {
	var bucket StorageBucket
	err := Connection.QueryRow("SELECT * FROM storage_buckets WHERE id = ?", id).Scan(&bucket.Id, &bucket.Name, &bucket.Type, &bucket.MaxSize, &bucket.RootPath, &bucket.Ip, &bucket.Port, &bucket.Username, &bucket.Password, &bucket.CreatedAt)
	if err != nil {
		return StorageBucket{}, err
	}

	if bucket.Password != "" {
		decrypted, err := encryption.Decrypt(bucket.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to decrypt password")
		}

		bucket.Password = decrypted
	}

	return bucket, nil
}
func GetBucketByName(name string) (StorageBucket, error) {
	var bucket StorageBucket
	err := Connection.QueryRow("SELECT * FROM storage_buckets WHERE name = ?", name).Scan(&bucket.Id, &bucket.Name, &bucket.Type, &bucket.MaxSize, &bucket.RootPath, &bucket.Ip, &bucket.Port, &bucket.Username, &bucket.Password, &bucket.CreatedAt)
	if err != nil {
		return StorageBucket{}, err
	}

	if bucket.Password != "" {
		decrypted, err := encryption.Decrypt(bucket.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to decrypt password")
		}

		bucket.Password = decrypted
	}

	return bucket, nil
}
func CreateBucket(bucket StorageBucket) error {
	_, err := Connection.Exec("INSERT INTO storage_buckets (name, type, max_size, root_path, ip, port, username, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", bucket.Name, bucket.Type, bucket.MaxSize, bucket.RootPath, bucket.Ip, bucket.Port, bucket.Username, bucket.Password)
	return err
}
func GetAvailableBucket(filesize int) (StorageBucket, error) {
	rows, err := Connection.Query(`
  SELECT b.id, b.name, b.type, b.max_size, b.root_path, IFNULL(b.ip, ''), IFNULL(b.port, 0), IFNULL(b.username, ''), IFNULL(b.password, ''), b.created_at, COALESCE(SUM(o.size), 0) + ? as used_size
  FROM storage_buckets b
  LEFT JOIN storage_objects o ON b.id = o.bucket_id
  GROUP BY b.id
  HAVING b.max_size >= used_size
`, filesize)
	if err != nil {
		return StorageBucket{}, err
	}
	defer rows.Close()

	var buckets []StorageBucket
	for rows.Next() {
		var bucket StorageBucket
		var usedSize int
		err = rows.Scan(&bucket.Id, &bucket.Name, &bucket.Type, &bucket.MaxSize, &bucket.RootPath, &bucket.Ip, &bucket.Port, &bucket.Username, &bucket.Password, &bucket.CreatedAt, &usedSize)
		if err != nil {
			return StorageBucket{}, err
		}
		if bucket.Password != "" {
			decrypted, err := encryption.Decrypt(bucket.Password)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to decrypt password")
			}

			bucket.Password = decrypted
		}

		buckets = append(buckets, bucket)
	}

	if len(buckets) == 0 {
		return StorageBucket{}, errors.New("No available bucket found")
	}

	return buckets[rand.Intn(len(buckets))], nil
}
