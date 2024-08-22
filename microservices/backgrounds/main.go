package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/microservices/backgrounds/database"
	"github.com/fernandovmedina/netflix-clone-react/microservices/backgrounds/handlers"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Printf("Error on microservices/backgrounds: %s\n", err.Error())
	}

	if err = godotenv.Load(); err != nil {
		log.Printf("Error on microservices/backgrounds: %s\n", err.Error())
	}

	var (
		serverPort string = os.Getenv("SERVER_PORT")
	)

	var router = http.NewServeMux()

	router.HandleFunc("/microservice/backgrounds", handlers.Backgrounds)
	router.HandleFunc("/microservice/background", handlers.Background)

	handler := handlers.CORS(router)
	handler = handlers.Middleware(handler)

	var server = http.Server{
		Addr:           serverPort,
		Handler:        handler,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Printf("Microservice running on %s", serverPort)

	log.Fatal(server.ListenAndServe())
}
