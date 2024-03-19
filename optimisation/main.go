package optimisation

import (
	"log"
	"time"
)

const maxWorkers = 10

var queue = make(map[string]string)
var currentWorkers = 0

// optimise optimises a file based on the file type
func optimise(path string, fileType string) {
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

// AddToQueue adds a file to the queue for optimisation
func AddToQueue(path string, fileType string) {
	queue[path] = fileType
}

// StartQueueThread starts a thread that will optimise files in the queue
func StartQueueThread() {
	go func() {
		for {
			if len(queue) > 0 && currentWorkers < maxWorkers {
				for path, fileType := range queue {
					delete(queue, path)
					optimise(path, fileType)
					break
				}
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}
