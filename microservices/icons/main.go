package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone/microservices/icons/handlers"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Printf("Error on microservices/icons: %v\n", err.Error())
	}

	var (
		serverAdd string = os.Getenv("SERVER_ADDRESS")
	)

	var mux *http.ServeMux = http.NewServeMux()

	mux.HandleFunc("/microservice/title", handlers.Title)
	mux.HandleFunc("/microservice/titles", handlers.Titles)

	var server http.Server = http.Server{
		Addr:           serverAdd,
		Handler:        mux,
		WriteTimeout:   20 * time.Second,
		ReadTimeout:    20 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Printf("Microservice running on %s\n", serverAdd)

	log.Fatal(server.ListenAndServe())
}
