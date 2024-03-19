package main

import (
	"github.com/nietthijmen/tavern/config"
	"github.com/nietthijmen/tavern/optimisation"
	"github.com/nietthijmen/tavern/prometheus"
	"github.com/nietthijmen/tavern/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	domain := config.ReadEnv("DOMAIN", "localhost")
	port := config.ReadEnv("PORT", "8080")
	secret := config.ReadEnv("SECRET", "secret")
	enablePrometheus := config.ReadEnv("ENABLE_PROMETHEUS", "false") == "true"

	log.Println("Starting queue server for optimisation")
	optimisation.StartQueueThread()

	// start a http server
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))
	http.HandleFunc("/"+secret, routes.Upload)

	if enablePrometheus {
		prometheus.RecordMetrics()
		http.Handle("/metrics", promhttp.Handler())
	}

	var listen string
	if strings.Contains(runtime.GOOS, "windows") {
		listen = "127.0.0.1:" + port
	} else {
		listen = ":" + port
	}

	log.Printf("Starting server on host: %s & port %s. Secret path = %s", domain, port, secret)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
