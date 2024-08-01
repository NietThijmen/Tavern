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

	DriverMap = map[string]Driver{}
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

func InitialiseDriverMap() {
	buckets, err := database.GetAllBuckets()
	if err != nil {
		log.Error().Err(err).Msg("Error getting all buckets")
		return
	}

	for _, bucket := range buckets {
		DriverMap[bucket.Name] = GetDriverFromBucket(bucket)
		err := DriverMap[bucket.Name].Connect()
		if err != nil {
			log.Error().Err(err).Msg("Error connecting to driver")
		}
	}
}

func GetInitialisedDriver(bucket database.StorageBucket) Driver {
	driver := DriverMap[bucket.Name]
	if driver == nil {
		driver = GetDriverFromBucket(bucket)
		err := driver.Connect()
		if err != nil {
			log.Error().Err(err).Msg("Error connecting to driver")
		}
		DriverMap[bucket.Name] = driver
	}

	if driver.Started() == false {
		log.Error().Msg("Driver not started")
		driver.Connect()
	}

	return driver
}
