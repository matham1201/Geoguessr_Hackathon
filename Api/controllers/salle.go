package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"main.go/models"
)

func AllSalle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
		salle.Nom = r.FormValue("name")
		salle.Coordonnees_x, _ = strconv.ParseFloat(r.FormValue("cordinnates_x"), 64)
		salle.Coordonnees_y, _ = strconv.ParseFloat(r.FormValue("cordinnates_y"), 64)
		salle.Etage, _ = strconv.Atoi(r.FormValue("floor"))
		salle.Disponibilite, _ = strconv.ParseBool(r.FormValue("disponibility"))

		fmt.Println(salle)

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
