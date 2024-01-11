package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"main.go/models"
)

// func AllSalle(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		salles := models.GetAllScore()
// 		json.NewEncoder(w).Encode(salles)
// 	default:
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 	}
// }

func Score(w http.ResponseWriter, r *http.Request) { // GET a Score by id
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		if r.URL.Path == "/score" {
			scores := models.GetAllScore()
			json.NewEncoder(w).Encode(scores)
			return
		}

		idStr := r.URL.Path[len("/score/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		score := models.GetScore(id)
		if score == (models.Score{}) {
			// afficher un message d'erreur 404 en json
			RespondWithError(w, http.StatusNotFound, "404", "Score not found")
			return
		}
		json.NewEncoder(w).Encode(score)
	case "POST":
		var score models.Score
		score.Name = r.FormValue("name")
		fmt.Println(score.Name)
		scoreInt, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "Invalid score", http.StatusBadRequest)
			return
		}
		score.Score = scoreInt

		fmt.Println(score)

		models.AddScore(score)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
