package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type NClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	jwt.RegisteredClaims
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type, send only JSON"))
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusUnsupportedMediaType,
				"body": map[string]interface{}{
					"error": "415 - Unsupported Media Type, send only JSON",
				},
			})
			return
		} else {
			var nclaims = new(NClaims)

			if err := json.NewDecoder(r.Body).Decode(&nclaims); err != nil {
				log.Printf("Error on microservices/icons: %s", err.Error())
			}

			jwt.NewWithClaims(jwt.SigningMethodES256, nclaims)
		}
	})
}
