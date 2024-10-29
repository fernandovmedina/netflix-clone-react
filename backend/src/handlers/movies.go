package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Function to create a new movie and insert it to the database
func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	var name = r.FormValue("name")
	var createdAt = getDateTime()
	file, header, err := r.FormFile("background")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	var duration = r.FormValue("duration")
	var yearReleased = r.FormValue("year")
	var description = r.FormValue("description")
	var videoFile = "url"
}

// Function to get the actual datetime
func getDateTime() string {
	// YYYY-MM-DD HH:MM:SS
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
