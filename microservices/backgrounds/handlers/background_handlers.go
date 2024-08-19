package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/fernandovmedina/netflix-clone-react/microservices/backgrounds/database"
)

type Bkg struct {
	Id   int    `json:"id_background"`
	Name string `json:"name"`
	Src  string `json:"source"`
}

type Any map[string]interface{}

func Backgrounds(w http.ResponseWriter, r *http.Request) {
	var bkgs = make([]Bkg, 0)

	rows, err := database.DB.Query("SELECT S.ID_SERIE, S.NAME, B.SOURCE FROM SERIES S, BACKGROUNDS B WHERE S.ID_BACKGROUND = B.ID_BACKGROUND")

	if err != nil {
		json.NewEncoder(w).Encode(Any{
			"status_code": http.StatusInternalServerError,
			"body": Any{
				"error": err.Error(),
			},
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var bkg = new(Bkg)

		if err = rows.Scan(&bkg.Id, &bkg.Name, &bkg.Src); err != nil {
			json.NewEncoder(w).Encode(Any{
				"status_code": http.StatusInternalServerError,
				"body": Any{
					"error": err.Error(),
				},
			})
			return
		}

		bkgs = append(bkgs, *bkg)
	}

	json.NewEncoder(w).Encode(Any{
		"status_code": http.StatusOK,
		"body":        bkgs,
	})
}

func Background(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")

	var bname = strings.ReplaceAll(name, "+", " ")

	rows, err := database.DB.Query("SELECT B.SOURCE FROM SERIES S, BACKGROUNDS B WHERE S.ID_BACKGROUND = B.ID_BACKGROUND AND S.NAME = ?", &bname)

	if err != nil {
		json.NewEncoder(w).Encode(Any{
			"status_code": http.StatusBadRequest,
			"body": Any{
				"text":  "check out if the name is the correct one",
				"error": err.Error(),
			},
		})
	}

	defer rows.Close()

	var src string

	for rows.Next() {
		if err = rows.Scan(&src); err != nil {
			json.NewEncoder(w).Encode(Any{
				"status_code": http.StatusInternalServerError,
				"body": Any{
					"error": err.Error(),
				},
			})
		}
	}

	_, err = os.ReadFile(src)

	if err != nil {
		json.NewEncoder(w).Encode(Any{
			"status_code": http.StatusNotFound,
			"body": Any{
				"text":  "file not found",
				"error": err.Error(),
			},
		})
	}

	http.ServeFile(w, r, src)
}
