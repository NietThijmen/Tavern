package storage

import (
	"github.com/google/uuid"
	"io"
	"os"
)

type LocalDriver struct {
	RootPath string
}

func (d *LocalDriver) Connect() error {
	return nil
}

func (d *LocalDriver) Disconnect() error {
	return nil
}

func (d *LocalDriver) Started() bool {
	return true
}

func (d *LocalDriver) UploadFile(file LocalFile) (RemoteFile, error) {
	uuidForPath := uuid.New()
	fullPath := d.RootPath + "/" + uuidForPath.String() + "/" + file.Name

	err := os.MkdirAll(d.RootPath+"/"+uuidForPath.String(), os.ModePerm)
	if err != nil {
		return RemoteFile{}, err
	}

	targetFile, err := os.Create(fullPath)
	if err != nil {
		println("Error creating file " + err.Error())
		return RemoteFile{}, err
	}

	defer targetFile.Close()
	for {
		var buffer = make([]byte, 8096)
		n, err := file.Reader.Read(buffer)
		if err != nil && err != io.EOF {
			return RemoteFile{}, err
		}

		if n == 0 {
			break
		}

		_, err = targetFile.Write(buffer[:n])
	}

	return RemoteFile{
		Name: file.Name,
		Size: file.Size,
		Path: uuidForPath.String() + "/" + file.Name,
	}, nil
}

func (d *LocalDriver) StreamFile(path string) (io.ReadCloser, error) {
	fullPath := d.RootPath + "/" + path
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (d *LocalDriver) DownloadFile(path string, targetPath string) error {
	remote := d.RootPath + "/" + path
	openRemote, err := os.Open(remote)
	if err != nil {
		return err
	}

	createLocal, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	defer openRemote.Close()
	defer createLocal.Close()

	_, err = io.Copy(createLocal, openRemote)
	if err != nil {
		return err
	}

	return nil
}

func (d *LocalDriver) DeleteFile(path string) error {
	fullPath := d.RootPath + "/" + path
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}

func (d *LocalDriver) ListFiles(path string) ([]string, error) {
	fullPath := d.RootPath + "/" + path
	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func (d *LocalDriver) GetFile(path string) ([]byte, error) {
	fullPath := d.RootPath + "/" + path
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return io.ReadAll(file)
}

func NewLocalDriver(config map[string]string) Driver {
	var driver Driver
	driver = &LocalDriver{
		RootPath: config["root_path"],
	}

	return driver
}
