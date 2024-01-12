package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"main.go/config"
	"main.go/models"
)

func AllSalle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		if r.URL.Path == "/salle" {
			salles := models.GetAllSalle()
			json.NewEncoder(w).Encode(salles)
			return
		}

		idStr := r.URL.Path[len("/salle/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		salle := models.GetSalle(id)
		if salle == (models.Salle{}) {
			RespondWithError(w, http.StatusNotFound, "404", "Salle not found")
			return
		}
		json.NewEncoder(w).Encode(salle)
	case "POST":
		var salle models.Salle
		salle.Id = config.GetLengthRoom() + 1
		salle.Nom = r.FormValue("name")
		salle.Coordonnees_x, _ = strconv.ParseFloat(r.FormValue("cordinnates_x"), 64)
		salle.Coordonnees_y, _ = strconv.ParseFloat(r.FormValue("cordinnates_y"), 64)
		salle.Etage, _ = strconv.Atoi(r.FormValue("floor"))
		salle.Disponibilite, _ = strconv.ParseBool(r.FormValue("disponibility"))
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Unable to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save the uploaded file to the server
		filePath := filepath.Join(uploadDirectory, handler.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to create the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Unable to copy file", http.StatusInternalServerError)
			return
		}

		salle.Photo = filePath

		fmt.Println(salle)
		//insert the new room in the database
		_, err = config.Db().Exec("INSERT INTO room (id, name, coordinate_x, coordinate_y, floor, disponibility, photo) VALUES (?, ?, ?, ?, ?, ?, ?)", salle.Id, salle.Nom, salle.Coordonnees_x, salle.Coordonnees_y, salle.Etage, salle.Disponibilite, salle.Photo)
		if err != nil {
			http.Error(w, "Unable to insert the room", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Room created successfully")

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, status int, errorType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{
		"message": message,
		"error":   errorType,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
