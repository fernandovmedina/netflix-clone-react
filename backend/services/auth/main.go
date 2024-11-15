package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/handlers"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/middlewares"
)

func main() {
	var err error

	if _, err = database.ConnMongoDB(); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	if err = godotenv.Load(".env"); err != nil {
		log.Printf("Error on main: %s\n", err.Error())
	}

	var (
		servicePort string = os.Getenv("SERVICE_PORT")
	)

	var mux = http.NewServeMux()

	mux.HandleFunc("POST /service/v1/auth/login", handlers.Login)
	mux.HandleFunc("POST /service/v1/auth/register", handlers.Register)
	mux.HandleFunc("GET /service/v1/auth/logout", handlers.Logout)

	var handler = middlewares.CORS(mux)

	var serviceRouter = http.Server{
		Addr:           servicePort,
		Handler:        handler,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Println("Auth service running on http://127.0.0.1:8010/service/v1/auth")

	log.Fatal(serviceRouter.ListenAndServe())
}
