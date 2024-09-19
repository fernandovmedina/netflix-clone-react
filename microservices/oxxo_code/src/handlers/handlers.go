package handlers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/joho/godotenv"

	"github.com/fernandovmedina/netflix-clone-react/microservices/oxxo/src/database"
)

func OxxoCode(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")
	var plan = r.URL.Query().Get("plan")

	if email == "" || plan == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusBadRequest,
			"body": map[string]any{
				"error": "email and plan are required",
			},
		})
		return
	}

	barCode, value, err := GenerateCode(email, plan)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	counter, err := GetLastCounter(email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	counter++

	route, err := CreateNewPngRoute(email, counter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	img, err := barcode.Scale(barCode, 400, 200)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	output, err := os.Create(route)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	defer output.Close()

	if err = png.Encode(output, img); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	_, returnedSource, err := InsertOxxoCode(value, route, email, counter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code": http.StatusInternalServerError,
			"body": map[string]any{
				"error": err.Error(),
			},
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": http.StatusCreated,
		"body": map[string]interface{}{
			"value":  value,
			"source": returnedSource,
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

func GetLastCounter(email string) (int, error) {
	var counter sql.NullInt16

	var err = database.DB.QueryRow("SELECT MAX(COUNTER) FROM OXXO_CODES WHERE EMAIL=?", email).Scan(&counter)

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if counter.Valid {
		return int(counter.Int16), nil
	}

	return 0, nil
}

func CreateNewPngRoute(email string, counter int) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	var route = os.Getenv("OUTPUT_ROUTE")

	if route == "" {
		route = "output"
	}

	name := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s@%d", email, counter)))

	path := filepath.Join(route, fmt.Sprintf("%s.png", name))

	return path, nil
}

func InsertOxxoCode(value, source, email string, counter int) (string, string, error) {
	var returnedValue, returnedSource string

	var err = database.DB.QueryRow("CALL InsertOxxoCode(?,?,?,?)", value, source, email, counter).Scan(&returnedValue, &returnedSource)

	if err != nil {
		return "", "", err
	}

	return returnedValue, returnedSource, nil
}
