package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	var (
		servicePort string = os.Getenv("SERVICE_PORT")
	)

	var mux = http.NewServeMux()

	var server = http.Server{
		Handler:        mux,
		Addr:           servicePort,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
