package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var loginRequest = new(models.LoginRequest)

		if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"error":       "Invalid request payload",
				"body":        nil,
			})
			return
		}

		var filter = bson.M{"email": loginRequest.Email}

		var user struct {
			ID             string `bson:"_id"`
			HashedPassword string `bson:"hashed_password"`
		}

		if err := database.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusNotFound,
				"error":       "User not found",
				"body":        nil,
			})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginRequest.HashedPassword)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusUnauthorized,
				"error":       "Invalid credentials",
				"body":        nil,
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Path:     "/",
			Value:    user.ID,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusOK,
			"message":     "Login successful",
			"body":        nil,
		})
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusMethodNotAllowed,
			"error":       "method not allowed",
			"body":        nil,
		})
		return
	}
}
