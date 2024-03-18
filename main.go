package main

import (
	"DiscordUpload/config"
	"DiscordUpload/routes"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	domain := config.ReadEnv("DOMAIN")
	port := config.ReadEnv("PORT")
	secret := config.ReadEnv("SECRET")

	log.Printf("Starting server on host: %s & port %s. Secret path = %s", domain, port, secret)

	// start a http server
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))
	http.HandleFunc("/"+secret, routes.Upload)

	var listen string
	if strings.Contains(runtime.GOOS, "windows") {
		listen = "0.0.0.0:" + port
	} else {
		listen = ":" + port
	}

	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
