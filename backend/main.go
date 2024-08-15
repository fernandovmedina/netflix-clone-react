package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/backend/src/database"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Println(err.Error())
	}

	var router = http.NewServeMux()

	var server = http.Server{
		Addr:         ":8000",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
