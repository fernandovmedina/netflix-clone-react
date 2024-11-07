package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/middlewares"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Printf("Erron on main: %s\n", err.Error())
	}

	var (
		servicePort string = os.Getenv("SERVICE_PORT")
	)

	var mux = http.NewServeMux()

	var handler = middlewares.CORS(mux)

	var serviceRouter = http.Server{
		Addr:           servicePort,
		Handler:        handler,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Fatal(serviceRouter.ListenAndServe())
}
