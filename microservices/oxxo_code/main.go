package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/microservices/oxxo/src/database"
	"github.com/fernandovmedina/netflix-clone-react/microservices/oxxo/src/handlers"
)

func main() {
	var err error

	if _, err = database.ConnDB(); err != nil {
		fmt.Printf("Error on microservices/oxxo_code: %s\n", err.Error())
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("GET /microservice/oxxo_code", handlers.OXXO)

	var handler = handlers.CORS(mux)

	var server = http.Server{
		Handler:        handler,
		Addr:           ":8050",
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	log.Println("[OXXO_CODE] Microservice running on http://127.0.0.1:8050")

	log.Fatal(server.ListenAndServe())
}
