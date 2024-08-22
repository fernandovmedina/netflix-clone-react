package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/database"
	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/handlers"
	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/middleware"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Printf("Error on microservices/auth: %s\n", err.Error())
	}

	if err = godotenv.Load(); err != nil {
		log.Printf("Error on microservices/auth: %s\n", err.Error())
	}

	var (
		serverPort string = os.Getenv("SERVER_PORT")
	)

	var mux = http.NewServeMux()

	mux.HandleFunc("/microservice/login", handlers.Login)
	mux.HandleFunc("/microservice/register", handlers.Register)
	mux.HandleFunc("/microservice/signout", handlers.SignOut)

	handler := middleware.CORS(mux)
	handler = middleware.Init(handler)

	var server = http.Server{
		Addr:           serverPort,
		Handler:        handler,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
