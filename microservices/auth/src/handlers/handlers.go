package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/fernandovmedina/netflix-clone-react/microservices/auth/src/database"
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
		Email    string            `json:"email"`
		Password string            `json:"password"`
		Plan     string            `json:"plan"`
		Users    map[string]User   `json:"users"`
		Payload  map[string]string `json:"payload"`
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

	var token = sha256.Sum256([]byte(register.Body.Email))

	var tokenSTR = hex.EncodeToString(token[:])

	passcrypt, err := bcrypt.GenerateFromPassword([]byte(register.Body.Password), bcrypt.DefaultCost)

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

	_, err = database.DB.Exec("INSERT INTO ACCOUNTS(EMAIL,PASS,TOKEN)VALUES(?,?,?)", register.Body.Email, passcrypt, tokenSTR)

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

	for _, v := range register.Body.Users {
		_, err = database.DB.Exec("INSERT INTO USERS(NAME,EMAIL,ID_ICON,ID_TYPE,ID_ACCOUNT)VALUES(?,?,?,?,(SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL=?))", v.Name, register.Body.Email, v.Icon, v.Type, register.Body.Email)

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

	switch {
	case register.Body.Payload["type"] == "gift":
		_, err = database.DB.Exec("CALL InsertPaymentAndGiftCode(?,?,?,?,?)", register.Body.Plan, time.Now(), register.Body.Email, "gift")

		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusExpectationFailed,
				"body": map[string]interface{}{
					"error": err.Error(),
				},
			})
			return
		}
	case register.Body.Payload["type"] == "oxxo":
		_, err = database.DB.Exec("CALL InsertPaymentAndOxxo(?,?,?,?,?,?,?)", register.Body.Plan, time.Now(), register.Body.Email, "oxxo", register.Body.Payload["phone_number"], register.Body.Payload["name"], register.Body.Payload["code"])

		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusExpectationFailed,
				"body": map[string]interface{}{
					"error": err.Error(),
				},
			})
			return
		}
	case register.Body.Payload["type"] == "card":
		_, err = database.DB.Exec("CALL InsertPaymentAndCard(?,?,?,?,?,?,?)", register.Body.Plan, time.Now(), register.Body.Email, "card", register.Body.Payload["card_number"], register.Body.Payload["due_date"], register.Body.Payload["cvv"], register.Body.Payload["card_name"])

		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusExpectationFailed,
				"body": map[string]interface{}{
					"error": err.Error(),
				},
			})
			return
		}
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body": map[string]interface{}{
			"text": "account registered successfully",
		},
	})
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
