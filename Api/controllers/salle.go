package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main.go/config"
	"main.go/models"
)

const uploadDirectory = "./uploads/"
const maxUploadSize = 10 * 1024 * 1024 // 10 MB

func AllSalle(w http.ResponseWriter, r *http.Request) {

	// Add CORS headers
	enableCors(&w)
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
		err := json.NewDecoder(r.Body).Decode(&salle)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid JSON payload")
			return
		}

		//insert the new room in the database
		_, err = config.Db().Exec("INSERT INTO room (id, name, coordinate_x, coordinate_y, floor, disponibility, photo) VALUES (?, ?, ?, ?, ?, ?, ?)", salle.Id, salle.Name, salle.Cordinnates_x, salle.Cordinnates_y, salle.Floor, salle.Disponibility, salle.Photo)
		if err != nil {
			http.Error(w, "Unable to insert the room", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case "PUT":
		idStr := r.URL.Path[len("/salle/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		var salle models.Salle
		json.NewDecoder(r.Body).Decode(&salle)
		salle.Id = id
		models.UpdateSalle(salle)
		json.NewEncoder(w).Encode(salle)
	case "DELETE":
		idStr := r.URL.Path[len("/salle/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		models.DeleteSalle(id)
		w.WriteHeader(http.StatusNoContent)
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-Type", "application/json")

}
