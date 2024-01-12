package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"main.go/config"
)

const uploadDirectory = "./uploads/"
const maxUploadSize = 10 * 1024 * 1024 // 10 MB

func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetImage")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/image/"):]
	fmt.Println(id)
	var image string

	err := config.Db().QueryRow("SELECT photo FROM room WHERE id = ?", id).Scan(&image)
	if err != nil {
		http.Error(w, "Unable to get the image", http.StatusInternalServerError)
		return
	}
	fmt.Println(image)

	http.ServeFile(w, r, image)
}
