package handlers

import (
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/fernandovmedina/netflix-clone-react/microservices/oxxo/src/database"
	"github.com/joho/godotenv"
)

func OxxoCode(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	var values = strings.Split(cookie.Value, " ")

	var email = values[0]

	_, value, err := GenerateCode(email)

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

	_, err = database.DB.Exec("INSERT INTO OXXO_CODES(VALUE,ID_ACCOUNT)VALUES(?,(SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL=?))", &value, &email)

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

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body": map[string]interface{}{
			"text": value,
		},
	})
}

func OxxoCodeIMG(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	var values = strings.Split(cookie.Value, " ")

	var email = values[0]

	barCode, _, err := GenerateCode(email)

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

	img, err := barcode.Scale(barCode, 300, 100)

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

	route, err := CreateNewPngRoute(email)

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

	output, err := os.Create(route)

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

	defer output.Close()

	if err = png.Encode(output, img); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status_code": http.StatusInternalServerError,
			"body": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	_, err = database.DB.Exec("UPDATE OXXO_CODES SET SOURCE=? WHERE ID_ACCOUNT=(SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL=?)", &route, &email)

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

	http.ServeFile(w, r, route)
}

func GenerateCode(email string) (barcode.BarcodeIntCS, string, error) {
	var uuid string
	var amount float32

	var err = database.DB.QueryRow("SELECT A.UUID, P.AMOUNT FROM ACCOUNTS A, PAYMENTS P WHERE P.ID_ACCOUNT=A.ID_ACCOUNT AND A.EMAIL=?", &email).Scan(&uuid, &amount)

	if err != nil {
		return nil, "", err
	}

	barCode, err := code128.Encode(fmt.Sprintf("%s %v", uuid, amount))

	if err != nil {
		return nil, "", err
	}

	return barCode, barCode.Content(), nil
}

func CreateNewPngRoute(email string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	var route = os.Getenv("OUTPUT_ROUTE")

	var uuid string

	if err := database.DB.QueryRow("SELECT UUID FROM ACCOUNTS WHERE EMAIL=?", &email).Scan(&uuid); err != nil {
		return "", err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.png", route, uuid))

	if err != nil {
		return "", err
	}

	return file.Name(), nil
}
