package main

import (
	"log"
	"net/http"

	"main.go/config"
)

func main() {
	config.DatabaseInit()
	// router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
