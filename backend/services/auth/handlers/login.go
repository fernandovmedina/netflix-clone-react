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
	var loginRequest = new(models.LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
			"body":        nil,
		})
		return
	}

	var filter = bson.M{
		"email": loginRequest.Email,
	}

	var result struct {
		_id             string `bson:"_id"`
		hashed_password string `bson:"hashed_password"`
	}

	if err := database.DB.Collection("users").FindOne(context.Background(), filter).Decode(&result); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusNotFound,
			"error":       "user not found",
			"body":        nil,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.hashed_password), []byte(loginRequest.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusUnauthorized,
			"error":       "invalid credentials",
			"body":        nil,
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Path:     "/",
		Value:    result._id,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusOK,
		"message":     "login successful",
		"body":        nil,
	})
}
