package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cookie, err := r.Cookie("auth")

		if err != nil {
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusNotImplemented,
				"error":       err.Error(),
				"body":        nil,
			})
			return
		}

		var cookieValue = cookie.Value

		id, err := primitive.ObjectIDFromHex(cookieValue)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"error":       err.Error(),
				"body":        nil,
			})
			return
		}

		var result bson.M

		if err = database.DB.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&result); err != nil {
			if err == mongo.ErrNoDocuments {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status_code": http.StatusNotFound,
					"error":       "user not found",
					"body":        nil,
				})
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status_code": http.StatusInternalServerError,
					"error":       err.Error(),
					"body":        nil,
				})
				return
			}
		}

		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusFound,
			"message":     "user found",
			"body":        &result,
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
