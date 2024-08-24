package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/database"
	"golang.org/x/crypto/bcrypt"
)

// Cookie.Value = example@example.com password

func Login(w http.ResponseWriter, r *http.Request) {
	var email = r.FormValue("email")
	var password = r.FormValue("password")

	var dbPassword string

	if err := database.DB.QueryRow("SELECT PASS FROM ACCOUNTS WHERE EMAIL=?", &email).Scan(&dbPassword); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err != nil {
		w.WriteHeader(http.StatusForbidden)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusNotFound,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Path:     "/",
			Value:    fmt.Sprintf("%s %s", email, dbPassword),
			Expires:  time.Now().Add(time.Minute * 60 * 24),
			HttpOnly: true,
		})

		w.WriteHeader(http.StatusAccepted)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusAccepted,
			"body": map[string]interface{}{
				"text": "user authenticathed",
			},
		})
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
		Type string `json:"type"`
	}

	type Body struct {
		Email    string          `json:"email"`
		Password string          `json:"password"`
		Users    map[string]User `json:"users"`
	}

	type RegisterRequest struct {
		Body Body `json:"body"`
	}

	var register RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	var users []User

	for _, v := range register.Body.Users {
		users = append(users, v)
	}

	_, err := database.DB.Exec("")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

}

func SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusConflict,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	cookie.Expires = time.Now().Add(1 * -time.Minute)
}