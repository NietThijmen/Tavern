package storage

import (
	"github.com/nietthijmen/tavern/src/database"
	"github.com/rs/zerolog/log"
	"strconv"
)

var (
	// DriverTypes Used for the CLI interface, to list all available driver types
	DriverTypes = []string{"local", "sftp"}
	// DriverFields Used for the CLI interface, to list all required fields for each driver type
	DriverFields = map[string][]string{
		"local": {"Name", "MaxSize", "RootPath"},
		"sftp":  {"Name", "MaxSize", "RootPath", "Ip", "Port", "Username", "Password"},
	}
)

func GetDriver(driverType string, config map[string]string) Driver {
	switch driverType {
	case "local":
		return NewLocalDriver(config)
	case "sftp":
		return NewSFTPDRiver(config)
	}

	log.Error().Msg("Driver not found for type: " + driverType)
	return nil
}

func GetDriverFromBucket(bucket database.StorageBucket) Driver {
	config := map[string]string{
		"root_path": bucket.RootPath,
		"ip":        bucket.Ip,
		"port":      strconv.Itoa(bucket.Port),
		"username":  bucket.Username,
		"password":  bucket.Password,
	}

	return GetDriver(bucket.Type, config)
}
