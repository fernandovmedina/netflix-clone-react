package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from root /"))
	})

	var handler = middlewares.CORS(mux)

	var serviceRouter = http.Server{
		Addr:           servicePort,
		Handler:        handler,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Println("Auth service running on http://127.0.0.1:8010/service/auth")

	log.Fatal(serviceRouter.ListenAndServe())
}
