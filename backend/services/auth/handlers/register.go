package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/database"
	"github.com/fernandovmedina/netflix-clone-react/backend/services/auth/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
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

	var currentDate = GetFormattedDate()

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

	var document = bson.D{
		{Key: "email", Value: RegisterRequest.Email},
		{Key: "hashed_password", Value: hashedPassword},
		{Key: "plan_details", Value: bson.D{
			{Key: "type", Value: RegisterRequest.Plan.Type},
			{Key: "device_limit", Value: RegisterRequest.Plan.DeviceLimit},
			{Key: "resolution", Value: RegisterRequest.Plan.Resolution},
		}},
		{Key: "users", Value: GetUsers(*RegisterRequest)},
		{Key: "payment", Value: bson.D{
			{Key: "card", Value: bson.D{
				{Key: "card_number", Value: RegisterRequest.Payment.CARD.Number},
				{Key: "due_date", Value: RegisterRequest.Payment.CARD.DueDate},
				{Key: "cvv", Value: RegisterRequest.Payment.CARD.CVV},
				{Key: "card_name", Value: RegisterRequest.Payment.CARD.Name},
			}},
		}},
		{Key: "initial_date", Value: currentDate},
		{Key: "expires_date", Value: AddDays(currentDate, 30)},
		{Key: "amount", Value: RegisterRequest.Amount},
		{Key: "status", Value: true},
		{Key: "auto_renew", Value: false},
	}

	result, err := database.DB.Collection("users").InsertOne(context.Background(), document)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
			"body":        nil,
		})
		return
	}

	var insertedID = result.InsertedID

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    fmt.Sprintf("%v", insertedID),
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body":        fmt.Sprintf("success, last inserted id %v", insertedID),
	})
}

// Function to get all the users from the request
func GetUsers(u models.RegisterRequest) interface{} {
	var users = bson.D{}

	for _, v := range u.Users {
		var user = bson.D{
			{Key: "name", Value: v.Name},
			{Key: "user_type", Value: v.UserType},
			{Key: "icon", Value: v.Icon},
		}

		users = append(users, user...)
	}

	return users
}

// Function to get the current date on database format (LINUX)
func GetFormattedDate() string {
	return time.Now().Format("2006-01-02")
}

/*
Function to add the days based on the user preference it is going to be
*/
func AddDays(dateSTR string, days int) string {
	date, err := time.Parse("2006-01-02", dateSTR)

	if err != nil {
		log.Printf("Error on ./handlers/register.go: %v\n", err.Error())
	}

	return date.AddDate(0, 0, days).Format("2004-12-30")
}
