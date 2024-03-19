package optimisation

import (
	"github.com/nietthijmen/tavern/prometheus"
	"image"
	"image/jpeg"
	"log"
	"os"
)

// optimiseJpeg optimises a jpeg image with the given compression level
func optimiseJpeg(path string, compressionLevel int) (bool, string) {
	var err error
	file, err := os.Open(path)
	if err != nil {
		return false, err.Error()
	}

	oldStat, _ := file.Stat()
	var img image.Image
	img, err = jpeg.Decode(file)

	if err != nil {
		return false, err.Error()
	}

	err = file.Close()
	if err != nil {
		return false, err.Error()
	}

	file, err = os.Create(path)
	if err != nil {
		return false, err.Error()
	}

	var options = jpeg.Options{
		Quality: 100 - compressionLevel*10,
	}

	err = jpeg.Encode(file, img, &options)
	if err != nil {
		return false, err.Error()
	}

	newStat, _ := file.Stat()
	err = file.Close()
	if err != nil {
		return false, err.Error()
	}
	log.Printf("Optimised Jpeg: %s, from %d to %d", path, oldStat.Size(), newStat.Size())

	prometheus.SavedSpace.Add(float64(oldStat.Size() - newStat.Size()))

	return true, ""
}
