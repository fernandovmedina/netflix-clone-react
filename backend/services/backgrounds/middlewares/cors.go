package middlewares

import "net/http"

/*
	-- date:   October 21, 23:22 (UTC-5)
	-- author: Fernando Vazquez
*/

// Function to allow petitions from anywhere
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		h.ServeHTTP(w, r)
	})
}
