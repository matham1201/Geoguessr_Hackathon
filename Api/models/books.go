package models

import (
	"log"

	"main.go/config"
)

// Book is a representation of a book

type Salle struct {
	Id            int    `json:"id"`
	Nom           string `json:"nom"`
	Photo         string `json:"photo"`
	Disponibilite bool   `json:"disponibilite"`
	Coordonnees_x string `json:"coordonnees_x"`
	Coordonnees_y string `json:"coordonnees_y"`
	Etage         int    `json:"etage"`
}

type Salles []Salle

type Admin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetLengthSalle() int {
	var length int
	err := config.Db().QueryRow("SELECT COUNT(*) FROM salle").Scan(&length)
	if err != nil {
		log.Fatal(err)
	}
	return length
}

func GetSalleById(id int) Salle {
	var s Salle
	err := config.Db().QueryRow("SELECT * FROM salle WHERE id = $1", id).Scan(&s.Id, &s.Nom, &s.Photo)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func GetSalleByEtage(etage int) Salles {
	var s Salles
	rows, err := config.Db().Query("SELECT * FROM salle WHERE etage = $1", etage)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var salles Salles // Change variable type to Salles
	for rows.Next() {
		var s Salle
		err := rows.Scan(&s.Id, &s.Nom, &s.Photo, &s.Disponibilite, &s.Coordonnees_x, &s.Coordonnees_y, &s.Etage)
		if err != nil {
			log.Fatal(err)
		}
		salles = append(salles, s) // Append to salles instead of s
	}
	return s
}

func CheckAdmin(login string, password string) bool {
	var a Admin
	err := config.Db().QueryRow("SELECT * FROM admin WHERE login = $1 AND password = $2", login, password).Scan(&a.Login, &a.Password)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
