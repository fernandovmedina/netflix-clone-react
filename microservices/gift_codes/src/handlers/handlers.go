package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/fernandovmedina/netflix-clone-react/microservices/gift/src/database"
)

type GiftCode struct {
	Value string `json:"gift_code"`
}

func New(w http.ResponseWriter, r *http.Request) {
	// Get all the gift codes values from the database
	rows, err := database.DB.Query("SELECT VALUE FROM GIFT_CODES")

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

	defer rows.Close()

	var codes []string

	for rows.Next() {
		var code string
		if err = rows.Scan(&code); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusInternalServerError,
				"body": map[string]interface{}{
					"error": err.Error(),
				},
			})
			return
		}
		codes = append(codes, code)
	}

	var code string

	for {
		code = Generate()
		if !Exists(code, codes) {
			break
		}
	}

	// Insert the new code into the database
	_, err = database.DB.Exec("INSERT INTO GIFT_CODES(VALUE)VALUES(?)", &code)

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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body":        code,
	})
}

func Exists(code string, codes []string) bool {
	for _, v := range codes {
		if v == code {
			return true
		}
	}
	return false
}

func Generate() string {
	var giftCode string

	for i := 0; i < 4; i++ {
		if i == 3 {
			giftCode += Four()
		} else {
			giftCode += Four() + "-"
		}
	}

	return giftCode
}

func Four() string {
	var final string
	var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < 4; i++ {
		var index = rand.Intn(len(letters))
		final += string(letters[index])
	}

	return final
}
