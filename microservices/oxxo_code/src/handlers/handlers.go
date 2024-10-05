package handlers

import "net/http"

func OXXO(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")
	var plan = r.URL.Query().Get("plan")

}
