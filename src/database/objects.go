package database

import "time"

func CreateObject(object StorageObject) error {
	_, err := Connection.Exec("INSERT INTO storage_objects (bucket_id, slug, size, file_type, file_path, uploaded_by) VALUES (?, ?, ?, ?, ?, ?)", object.Bucket.Id, object.Slug, object.Size, object.FileType, object.FilePath, object.UploadedBy)
	if err != nil {
		return err
	}

	return nil
}
func FindObjectBySlug(slug string) (StorageObject, error) {
	var object StorageObject

	if cached, found := Cache.Get("object_" + slug); found {
		object = cached.(StorageObject)
		return object, nil
	}

	var bucketId int

	err := Connection.QueryRow("SELECT * FROM storage_objects WHERE slug = ?", slug).Scan(&object.Id, &bucketId, &object.Slug, &object.Size, &object.FileType, &object.FilePath, &object.UploadedBy, &object.CreatedAt)
	if err != nil {
		return object, err
	}

	object.Bucket, err = GetBucketById(bucketId)

	Cache.Set("object_"+slug, object, 1*time.Minute)

	return object, nil
}
