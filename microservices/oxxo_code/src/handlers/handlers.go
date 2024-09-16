package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/microservices/oxxo/src/database"
)

func OxxoCode(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")

	barCode, value, err := GenerateCode(email, r.URL.Query().Get("plan"))

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

	img, err := barcode.Scale(barCode, 400, 200)

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

	type OxxoCode struct {
		Value  string `json:"value"`
		Source string `json:"source"`
	}

	var oxxo OxxoCode

	var returnedValue string
	var returnedSource string

	err = database.DB.QueryRow("CALL InsertOxxoCode(?,?)", &value, &route).Scan(&returnedValue, &returnedSource)

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

	oxxo.Value = returnedValue
	oxxo.Source = returnedSource

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body": map[string]interface{}{
			"value":  oxxo.Value,
			"source": oxxo.Source,
		},
	})
}

func GenerateCode(email, plan string) (barcode.BarcodeIntCS, string, error) {
	barCode, err := code128.Encode(fmt.Sprintf("%s %s", email, plan))

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

	var counters []int

	rows, err := database.DB.Query("SELECT COUNTER FROM OXXO_CODES WHERE EMAIL=?", &email)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var counter int

		if err = rows.Scan(&counter); err != nil {
			return "", err
		}

		counters = append(counters, counter)
	}

	var counter int = GetBigger(counters)

	var name = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/%s@%d", route, email, counter)))

	file, err := os.Create(fmt.Sprintf("%s/%s.png", route, name))

	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func GetBigger(x []int) int {
	var biggest int = 0

	for _, v := range x {
		if v > biggest {
			biggest = v
		}
	}

	return biggest
}
