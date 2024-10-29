package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/backend/src/database"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if _, err = database.ConnMYSQL(); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	if err = godotenv.Load(".env"); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	var (
		serverPort string = os.Getenv("SERVER_PORT")
	)

	var mux = http.NewServeMux()

	var server = http.Server{
		Handler:        mux,
		Addr:           serverPort,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
