package controllers

import (
	"fmt"
	"net/http"

	"main.go/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var login models.Admin
		login.Username = r.FormValue("username")
		fmt.Println(login.Username)
		loginInt := r.FormValue("password")
		fmt.Println(loginInt)

		if models.CheckConnection(login.Username, loginInt) {
			fmt.Println("ok")
		} else {
			fmt.Println("pas ok")
		}

		fmt.Println(login)
	}
}
