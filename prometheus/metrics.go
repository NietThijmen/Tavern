package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"os"
	"time"
)

func RecordMetrics() {
	go func() {
		for {

			dirs, err := os.ReadDir("storage")
			if err != nil {
				continue
			}

			itemsUploaded.Set(float64(len(dirs)) - 1)

			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	itemsUploaded = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "items_uploaded",
		Help: "The total number of items uploaded",
	})
)
