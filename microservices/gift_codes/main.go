package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/microservices/gift/src/database"
	"github.com/fernandovmedina/netflix-clone-react/microservices/gift/src/handlers"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		log.Printf("Error on microservices/gift_codes: %s\n", err.Error())
	}

	if err = godotenv.Load(); err != nil {
		log.Printf("Error on microservices/gift_codes: %s\n", err.Error())
	}

	var (
		routerPort string = os.Getenv("ROUTER_PORT")
	)

	var mux = http.NewServeMux()

	mux.HandleFunc("GET /microservice/gift_code", handlers.New)

	var server = http.Server{
		Handler:        mux,
		Addr:           routerPort,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

  log.Println("[GIFT_CODES] Microservice running on http://127.0.0.1:8040/microservice/gift_code")

	log.Fatal(server.ListenAndServe())
}
