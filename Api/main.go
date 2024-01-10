package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/config"
	"main.go/controllers"
)

func main() {
	config.DatabaseInit()
	http.HandleFunc("/score", controllers.Score)    // GET & Post
	http.HandleFunc("/score/", controllers.Score)   //Get by Id
	http.HandleFunc("/salle", controllers.AllSalle) // GET & Post
	http.HandleFunc("/salle/", controllers.AllSalle)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
