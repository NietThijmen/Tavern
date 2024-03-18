package main

import (
	"DiscordUpload/config"
	"DiscordUpload/routes"
	"log"
	"net/http"
)

func main() {
	domain := config.ReadEnv("DOMAIN")
	port := config.ReadEnv("PORT")
	secret := config.ReadEnv("SECRET")

	log.Printf("Starting server on host: %s & port %s. Secret path = %s", domain, port, secret)

	// start a http server
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))
	http.HandleFunc("/"+secret, routes.Upload)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
