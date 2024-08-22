package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/database"
)

func Init(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusUnsupportedMediaType,
				"body": map[string]interface{}{
					"error": "415 - Unsupported Media Type, only JSON",
				},
			})
			return
		}

		cookie, err := r.Cookie("auth")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusNotFound,
				"body": map[string]interface{}{
					"error": "404 - Cookie not found on http request",
				},
			})
			return
		}

		rawValues := strings.SplitN(cookie.Value, " ", 2)

		if len(rawValues) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"body": map[string]interface{}{
					"error": "400 - Invalid cookie format",
				},
			})
			return
		}

		email, providedToken := rawValues[0], rawValues[1]

		var storedToken string

		err = database.DB.QueryRow("SELECT TOKEN FROM ACCOUNTS WHERE EMAIL = ?", email).Scan(&storedToken)

		if err != nil || storedToken != providedToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"body": map[string]interface{}{
					"error": "401 - Unauthorized, token mismatch",
				},
			})
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
