package database

type User struct {
	Id        int
	Email     string
	Password  string
	CreatedAt string
}

type ApiKey struct {
	Id        int
	User      User // This should be mapped to the User struct, not just the user_id.
	ApiKey    string
	CreatedAt string
}

type StorageBucket struct {
	Id        int
	Name      string
	Type      string
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
	Name       string
	Size       int
	FileType   string
	FilePath   string
	UploadedBy string
	CreatedAt  string
}
