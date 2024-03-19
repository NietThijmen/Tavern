package optimisation

import (
	"log"
	"time"
)

var maxWorkers = 10

var queue = make(map[string]string, 0)
var currentWorkers = 0

func Optimise(path string, fileType string) {
	switch fileType {
	case "image/png":
		currentWorkers++
		optimised, err := optimisePng(path, 5)
		currentWorkers--
		if !optimised {
			log.Printf("Failed to optimise PNG: %s", err)
		}
	case "image/jpeg":
		currentWorkers++
		optimised, err := optimiseJpeg(path, 5)
		currentWorkers--

		if !optimised {
			log.Printf("Failed to optimise JPEG: %s", err)
		}
	default:
		log.Printf("Unknown file type: %s", fileType)
	}
}

func AddToQueue(path string, fileType string) {
	queue[path] = fileType
}

func StartQueueThread() {
	go func() {
		for {
			if len(queue) > 0 && currentWorkers < maxWorkers {
				for path, fileType := range queue {
					delete(queue, path)
					Optimise(path, fileType)
					break
				}
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}
