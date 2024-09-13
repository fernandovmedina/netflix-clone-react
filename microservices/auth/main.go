package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/database"
	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/handlers"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Printf("Error on microservices/auth: %s\n", err.Error())
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("/microservice/login", handlers.Login)
	mux.HandleFunc("/microservice/register", handlers.Register)
	mux.HandleFunc("/microservice/signout", handlers.SignOut)

	var handler = handlers.CORS(mux)

	var server = http.Server{
		Addr:           ":8030",
		Handler:        handler,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Println("[AUTH] Microservice running on http://127.0.0.1:8030")

	log.Fatal(server.ListenAndServe())
}
