package main

/*
	-- date:   October 21, 23:16 (UTC-5)
	-- author: Fernando Vazquez
*/

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/backgrounds/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/backgrounds/middlewares"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	if err = godotenv.Load(".env"); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	var (
		servicePort string = os.Getenv("SERVICE_PORT")
	)

	var mux = http.NewServeMux()

	var handler = middlewares.CORS(mux)

	var server = http.Server{
		Handler:        handler,
		Addr:           servicePort,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
