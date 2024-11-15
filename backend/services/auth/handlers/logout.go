package handlers

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false,
	})

	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
