package optimisation

import (
	"github.com/nietthijmen/tavern/prometheus"
	compression "github.com/nurlantulemisov/imagecompression"
	"image"
	"image/png"
	"log"
	"os"
)

func optimisePng(path string, compressionLevel int) (bool, string) {
	var err error
	file, err := os.Open(path)
	if err != nil {
		return false, err.Error()
	}

	oldStat, _ := file.Stat()
	var img image.Image
	img, err = png.Decode(file)

	if err != nil {
		return false, err.Error()
	}

	file.Close()

	compressing, _ := compression.New(compressionLevel)
	compressingImage := compressing.Compress(img)

	file, err = os.Create(path)
	if err != nil {
		return false, err.Error()
	}

	err = png.Encode(file, compressingImage)
	if err != nil {
		return false, err.Error()
	}

	newStat, _ := file.Stat()
	file.Close()

	log.Printf("Optimised PNG: %s, from %d to %d", path, oldStat.Size(), newStat.Size())

	prometheus.SavedSpace.Add(float64(oldStat.Size() - newStat.Size()))

	return true, ""
}
