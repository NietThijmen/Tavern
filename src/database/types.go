package database

type ApiKey struct {
	Id        int
	ApiKey    string
	CreatedAt string
}

type StorageBucket struct {
	Id        int
	Name      string
	Type      string
	MaxSize   int64
	RootPath  string
	Ip        string
	Port      int
	Username  string
	Password  string
	CreatedAt string
}

type StorageObject struct {
	Id         int
	Bucket     StorageBucket // This should be mapped to the StorageBucket struct, not just the bucket_id.
	Slug       string
	Size       int
	FileType   string
	FilePath   string
	UploadedBy string
	CreatedAt  string
}
