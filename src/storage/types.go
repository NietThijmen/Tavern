package storage

import "io"

type LocalFile struct {
	Name   string
	Size   int
	Reader io.Reader
}

type RemoteFile struct {
	Name string
	Size int
	Path string
}

type Driver interface {
	Connect() error
	Disconnect() error
	UploadFile(file LocalFile) (RemoteFile, error)
	DownloadFile(path string, targetPath string) error
	DeleteFile(path string) error
	ListFiles(path string) ([]string, error)
	GetFile(path string) ([]byte, error)
}
