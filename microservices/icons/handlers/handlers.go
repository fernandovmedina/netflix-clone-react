package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fernandovmedina/netflix-clone/microservices/icons/database"
	"github.com/joho/godotenv"
)

func Title(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	} else {
		var err error

		if err = godotenv.Load(); err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}

		var id = r.URL.Query().Get("id")
		var title = r.URL.Query().Get("title")

		dir, err := os.ReadDir(os.Getenv("NETFLIX_IMAGES_FOLDER") + "/" + title)

		if err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}

		idINT, err := strconv.Atoi(id)

		if err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}

		if idINT > len(dir) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusBadRequest,
				"body": map[string]interface{}{
					"error": "id out of range",
				},
			})
			return
		} else {
			var imgURL = os.Getenv("NETFLIX_IMAGES_FOLDER") + "/" + title + "/" + dir[idINT-1].Name()

			fmt.Println(imgURL)

			var idTitle int

			if err := database.DB.QueryRow("SELECT ID_ICON FROM ICONS WHERE SOURCE=?", imgURL).Scan(&idTitle); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status_code": http.StatusNotFound,
					"body": map[string]interface{}{
						"error": err.Error(),
					},
				})
				return
			}

			w.WriteHeader(http.StatusFound)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"status_code": http.StatusFound,
				"body": map[string]interface{}{
					"id":     idTitle,
					"source": imgURL,
				},
			})
		}
	}
}

func Titles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	} else {
		var err error

		if err = godotenv.Load(); err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}

		directories, err := os.ReadDir(os.Getenv("NETFLIX_IMAGES_FOLDER"))

		if err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}

		type Dir struct {
			Name string
			Len  int
		}

		var dirs []Dir

		for i := range directories {
			dir, err := os.ReadDir(os.Getenv("NETFLIX_IMAGES_FOLDER") + "/" + directories[i].Name())

			if err != nil {
				log.Printf("Error on microservices/icons: %v\n", err.Error())
				return
			}

			dirs = append(dirs, Dir{Name: directories[i].Name(), Len: len(dir)})
		}

		if err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"body":   dirs,
		}); err != nil {
			log.Printf("Error on microservices/icons: %v\n", err.Error())
			return
		}
	}
}
