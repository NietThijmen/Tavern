package database

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

		buckets = append(buckets, bucket)
	}

	return buckets, nil
}
