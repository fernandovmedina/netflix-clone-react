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

	handler := handlers.CORS(mux)
	handler = handlers.Middleware(handler)

	var server http.Server = http.Server{
		Addr:           serverAdd,
		Handler:        handler,
		WriteTimeout:   20 * time.Second,
		ReadTimeout:    20 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

  log.Println("[ICONS] Microservice running on http://127.0.0.1:8010")

	log.Fatal(server.ListenAndServe())
}
