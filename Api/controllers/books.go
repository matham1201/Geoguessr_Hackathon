package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"main.go/models"
)

func SalleLength(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	salles := models.GetLengthSalle()
	if err := json.NewEncoder(w).Encode(salles); err != nil {
		log.Fatal(err)
	}
}
