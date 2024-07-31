package storage

import (
	"github.com/rs/zerolog/log"
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
