package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var RegisterRequest = new(models.RegisterRequest)

		if err := json.NewDecoder(r.Body).Decode(RegisterRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"error":       "bad request",
				"body":        nil,
			})
			return
		}

		if RegisterRequest.Email == "" || RegisterRequest.HashedPassword == "" || RegisterRequest.Plan.Type == "" || len(RegisterRequest.Payments) == 0 || len(RegisterRequest.Users) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"error":       "bad request, check out fields",
				"body":        nil,
			})
			return
		}

		var filter = bson.M{
			"email": RegisterRequest.Email,
		}

		var resultQuery struct{ _id string }

		err := database.DB.Collection("users").FindOne(context.Background(), filter).Decode(&resultQuery)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"error":       err.Error(),
				"body":        nil,
			})
			return
		}

		if err == nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusConflict,
				"message":     "user already exists",
				"body":        nil,
			})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegisterRequest.HashedPassword), bcrypt.DefaultCost)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"error":       err.Error(),
				"body":        nil,
			})
			return
		}
		var currentDate = GetFormattedDate()
		var expiresDate = AddDays(currentDate, 30)

		var document = bson.D{
			{Key: "email", Value: RegisterRequest.Email},
			{Key: "hashed_password", Value: hashedPassword},
			{Key: "plan_details", Value: bson.D{
				{Key: "type", Value: RegisterRequest.Plan.Type},
				{Key: "device_limit", Value: RegisterRequest.Plan.DeviceLimit},
				{Key: "resolution", Value: RegisterRequest.Plan.Resolution},
			}},
			{Key: "users", Value: GetUsers(*RegisterRequest)},
			{Key: "payment", Value: GetPayments(RegisterRequest.Payments)},
			{Key: "initial_date", Value: currentDate},
			{Key: "expires_date", Value: expiresDate},
			{Key: "amount", Value: RegisterRequest.Amount},
			{Key: "status", Value: true},
			{Key: "auto_renew", Value: false},
		}

		if _, err = database.DB.Collection("users").InsertOne(context.Background(), document); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"error":       err.Error(),
				"body":        nil,
			})
			return
		}

		filter = bson.M{"email": RegisterRequest.Email}

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

		http.SetCookie(w, &http.Cookie{
			Name:     "auth",
			Value:    user.ID,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
		})

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusCreated,
			"body":        fmt.Sprintf("success, last inserted id %v", user.ID),
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

// Function to get all the users from the request
func GetUsers(u models.RegisterRequest) interface{} {
	var users []bson.D

	for _, v := range u.Users {
		users = append(users, bson.D{
			{Key: "name", Value: v.Name},
			{Key: "user_type", Value: v.UserType},
			{Key: "icon", Value: v.Icon},
		})
	}

	return users
}

// Function to get the current date on database format (LINUX)
func GetFormattedDate() string {
	return time.Now().Format("2006-01-02")
}

// Function to add the days based on the user preference
func AddDays(dateSTR string, days int) string {
	date, err := time.Parse("2006-01-02", dateSTR)
	if err != nil {
		log.Printf("Error on ./handlers/register.go: %v\n", err.Error())
		return ""
	}

	return date.AddDate(0, 0, days).Format("2006-01-02")
}

func GetPayments(payments []models.Payment) interface{} {
	var paymentArray []bson.D

	for _, payment := range payments {
		if payment.Card != nil {
			paymentArray = append(paymentArray, bson.D{
				{Key: "card_number", Value: payment.Card.Number},
				{Key: "due_date", Value: payment.Card.DueDate},
				{Key: "cvv", Value: payment.Card.CVV},
				{Key: "card_name", Value: payment.Card.Name},
			})
		}
		if payment.Oxxo != nil {
			paymentArray = append(paymentArray, bson.D{
				{Key: "oxxo_name", Value: payment.Oxxo.Name},
				{Key: "oxxo_code", Value: payment.Oxxo.Code},
			})
		}
		if payment.Gift != nil {
			paymentArray = append(paymentArray, bson.D{
				{Key: "gift_code", Value: payment.Gift.Code},
			})
		}
	}

	return paymentArray
}
