package storage

import (
	"github.com/google/uuid"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"strconv"
	"time"
)

type SftpDriver struct {
	Username   string
	Password   string
	Ip         string
	Port       int
	RootPath   string
	Connection *sftp.Client
	IsStarted  bool
}

func (d *SftpDriver) Connect() error {
	config := &ssh.ClientConfig{
		User: d.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(d.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	addr := d.Ip + ":" + strconv.Itoa(d.Port)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}

	d.Connection = client
	d.IsStarted = true
	return nil
}

func (d *SftpDriver) Disconnect() error {
	d.IsStarted = false
	return d.Connection.Close()
}

func (d *SftpDriver) Started() bool {
	return d.IsStarted
}

func (d *SftpDriver) UploadFile(file LocalFile) (RemoteFile, error) {
	uuidForPath := uuid.New()
	fullPath := d.RootPath + "/" + uuidForPath.String() + "/" + file.Name

	err := d.Connection.MkdirAll(d.RootPath + "/" + uuidForPath.String())
	if err != nil {
		return RemoteFile{}, err
	}

	targetFile, err := d.Connection.Create(fullPath)
	if err != nil {
		println("Error creating file " + err.Error())
		return RemoteFile{}, err
	}

	defer targetFile.Close()

	for {
		var buffer = make([]byte, 1024)
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

func (d *SftpDriver) StreamFile(path string) (io.ReadCloser, error) {
	fullPath := d.RootPath + "/" + path
	file, err := d.Connection.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (d *SftpDriver) DownloadFile(path string, targetPath string) error {
	remote := d.RootPath + "/" + path
	openRemote, err := d.Connection.Open(remote)
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

func (d *SftpDriver) DeleteFile(path string) error {
	fullPath := d.RootPath + "/" + path
	err := d.Connection.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}

func (d *SftpDriver) ListFiles(path string) ([]string, error) {
	fullPath := d.RootPath + "/" + path
	files, err := d.Connection.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func (d *SftpDriver) GetFile(path string) ([]byte, error) {
	fullPath := d.RootPath + "/" + path
	file, err := d.Connection.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return io.ReadAll(file)
}

func NewSFTPDRiver(config map[string]string) Driver {
	var driver Driver

	PortAsInt, _ := strconv.Atoi(config["port"])

	driver = &SftpDriver{
		Username:   config["username"],
		Password:   config["password"],
		Ip:         config["ip"],
		Port:       PortAsInt,
		RootPath:   config["root_path"],
		Connection: &sftp.Client{},
	}

	return driver
}
