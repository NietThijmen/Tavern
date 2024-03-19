package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"os"
	"path/filepath"
	"time"
)

// RecordMetrics records metrics for the storage folder
func RecordMetrics() {
	go func() {
		for {

			dirs, err := os.ReadDir("storage")
			if err == nil {
				itemsUploaded.Set(float64(len(dirs)) - 1)
			}

			var totalSize int64
			err = filepath.Walk("storage", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				totalSize += info.Size()
				return nil
			})

			if err == nil {
				TotalSize.Set(float64(totalSize))
			}

			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	itemsUploaded = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "items_uploaded",
		Help: "The total number of items uploaded",
	})

	TotalSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_size",
		Help: "The total size of the storage folder",
	})

	SavedSpace = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "saved_space",
		Help: "The total space saved by optimising images",
	})
)
