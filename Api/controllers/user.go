package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "POST":
		var login models.Admin
		json.NewDecoder(r.Body).Decode(&login)
		fmt.Println("user", login.Username)
		fmt.Println("password", login.Password)

		if models.CheckConnection(login.Username, login.Password) {
			RespondWithJSON(w, http.StatusOK, "ok")
		} else {
			RespondWithJSON(w, http.StatusNotFound, "error")
		}

	}
}
