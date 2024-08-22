package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/backend/src/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/src/middleware"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Println(err.Error())
	}

	if err = godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	var (
		serverPort string = os.Getenv("SERVER_PORT")
	)

	var router = http.NewServeMux()

	handler := middleware.CORS(router)
	handler = middleware.Init(handler)

	var server = http.Server{
		Addr:           serverPort,
		Handler:        handler,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
