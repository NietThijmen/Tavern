package main

import (
	"DiscordUpload/config"
	"DiscordUpload/prometheus"
	"DiscordUpload/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	domain := config.ReadEnv("DOMAIN")
	port := config.ReadEnv("PORT")
	secret := config.ReadEnv("SECRET")
	enablePrometheus := config.ReadEnv("ENABLE_PROMETHEUS") == "true"

	log.Printf("Starting server on host: %s & port %s. Secret path = %s", domain, port, secret)

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

	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
