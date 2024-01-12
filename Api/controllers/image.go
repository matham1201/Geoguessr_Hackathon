package controllers

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"main.go/config"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/image/"):]
	var image string

	err := config.Db().QueryRow("SELECT photo FROM room WHERE id = ?", id).Scan(&image)
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, image)
}
