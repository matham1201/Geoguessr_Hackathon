package main

import (
	"github.com/gorilla/mux"
	"main.go/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	//router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.CarsCreate)
	//router.Methods("GET").Path("/cars").Name("Index").HandlerFunc(controllers.CarsIndex)
	//router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
	//router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
	//router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.CarsDelete)

	//Get /Jeux --> recupere toutes les infos des salle --> nom, photo et coor
	//Get /Jeux/Start --> recupere 5 salles random
	//Get /Jeux/Salle/{id} --> recupere la salle {id} --> nom, photo et coor

	//Get /Visite/salle --> toutes les salles
	//Get /Visite/salle/{id} --> salle {id} , nom et disponibilité

	//Post /Visite/salle/{id} --> salle {id} , nom et disponibilité
	//Post /Visite/salle/{id}/reserver --> salle {id} , nom et disponibilité

	router.Methods("GET").Path("/length").Name("taille").HandlerFunc(controllers.SalleLength)

	return router
}
