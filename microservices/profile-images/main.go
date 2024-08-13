package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone/microservices/profile-images/handlers"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Printf("Error on microservices/profile-images: %v\n", err.Error())
	}

	var (
		serverAdd string = os.Getenv("SERVER_ADDRESS")
	)

	var mux *http.ServeMux = http.NewServeMux()

	mux.HandleFunc("/microservice/title", handlers.Title)
	mux.HandleFunc("/microservice/titles", handlers.Titles)

	var server http.Server = http.Server{
		Addr:           serverAdd,
		Handler:        Middleware(mux),
		WriteTimeout:   20 * time.Second,
		ReadTimeout:    20 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Printf("Microservice running on %s\n", serverAdd)

	log.Fatal(server.ListenAndServe())
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configura los encabezados CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si es una solicitud OPTIONS, responde de inmediato
		if r.Method == http.MethodOptions {
			return
		}

		// Llama al siguiente handler
		next.ServeHTTP(w, r)
	})
}
