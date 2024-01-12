package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"main.go/config"
	"main.go/models"
)

const uploadDirectory = "./uploads/"
const maxUploadSize = 10 * 1024 * 1024 // 10 MB

func Salle(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	enableCors(&w) // On autorise les requêtes cross-origin
	switch r.Method {
	case "GET": // Method GET
		if r.URL.Path == "/salle" { // Si l'url est /salle
			salles := models.GetAllSalle()
			json.NewEncoder(w).Encode(salles) // On encode les salles en json
			return
		}

		idStr := r.URL.Path[len("/salle/"):] // On récupère l'id de la salle dans l'url
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		salle := models.GetSalle(id) // On récupère la salle dans la base de données
		if salle == (models.Salle{}) {
			RespondWithError(w, http.StatusNotFound, "404", "Salle not found")
			return
		}
		json.NewEncoder(w).Encode(salle) // On encode la salle en json
	case "POST": // Method POST
		var salle models.Salle // On crée une nouvelle salle
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
		dst, err := os.Create(filePath) // On crée le fichier
		if err != nil {
			http.Error(w, "Unable to create the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil { // On copie le fichier
			http.Error(w, "Unable to copy file", http.StatusInternalServerError)
			return
		}

		salle.Photo = filePath // On récupère le chemin du fichier

		// On insère la salle dans la base de données
		_, err = config.Db().Exec("INSERT INTO room (id, name, coordinate_x, coordinate_y, floor, disponibility, photo) VALUES (?, ?, ?, ?, ?, ?, ?)", salle.Id, salle.Nom, salle.Coordonnees_x, salle.Coordonnees_y, salle.Etage, salle.Disponibilite, salle.Photo)
		if err != nil {
			http.Error(w, "Unable to insert the room", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated) // On renvoie le code 201

	case "PUT": // Method PUT
		idStr := r.URL.Path[len("/salle/"):] // On récupère l'id de la salle dans l'url
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		var salle models.Salle
		json.NewDecoder(r.Body).Decode(&salle) // On décode la salle en json
		salle.Id = id
		models.UpdateSalle(salle)        // On met à jour la salle dans la base de données
		json.NewEncoder(w).Encode(salle) // On encode la salle en json
	case "DELETE": // Method DELETE
		idStr := r.URL.Path[len("/salle/"):] // On récupère l'id de la salle dans l'url
		id, err := strconv.Atoi(idStr)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "400", "Invalid ID")
			return
		}

		models.DeleteSalle(id)              // On supprime la salle de la base de données
		w.WriteHeader(http.StatusNoContent) // On renvoie le code 204
	case "OPTIONS":
		w.WriteHeader(http.StatusOK) // On renvoie le code 200

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed) // On renvoie le code 405
	}
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) { // On renvoie le payload en json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, status int, errorType, message string) { // On renvoie une erreur en json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{
		"message": message,
		"error":   errorType,
	}
	json.NewEncoder(w).Encode(errorResponse)
}

func enableCors(w *http.ResponseWriter) { // On autorise les requêtes cross-origin
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-Type", "application/json")

}
