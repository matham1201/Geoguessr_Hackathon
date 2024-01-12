package models

import (
	"log"

	"main.go/config"
)

// Book is a representation of a book

type Salle struct {
	Id            int     `json:"id"`
	Nom           string  `json:"name"`
	Coordonnees_x float64 `json:"cordinnates_x"`
	Coordonnees_y float64 `json:"cordinnates_y"`
	Etage         int     `json:"floor"`
	Disponibilite bool    `json:"disponibility"`
	Photo         string  `json:"photo"`
}

type Salles []Salle

func GetAllSalle() Salles {

	var salles Salles

	rows, err := config.Db().Query("SELECT id, name, coordinate_x, coordinate_y, floor, disponibility, photo FROM room")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var salle Salle
		if err := rows.Scan(&salle.Id, &salle.Nom, &salle.Coordonnees_x, &salle.Coordonnees_y, &salle.Etage, &salle.Disponibilite, &salle.Photo); err != nil {
			log.Fatal(err)
		}

		salles = append(salles, salle)
	}
	return salles
}

func GetSalle(id int) Salle {

	var salle Salle

	rows, err := config.Db().Query("SELECT id, name, coordinate_x, coordinate_y, floor, disponibility, photo FROM room WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&salle.Id, &salle.Nom, &salle.Coordonnees_x, &salle.Coordonnees_y, &salle.Etage, &salle.Disponibilite, &salle.Photo); err != nil {
			log.Fatal(err)
		}
	}
	return salle
}

func AddSalle(salle Salle) {
	//ceci est un commentaire
	stmt, err := config.Db().Prepare("INSERT INTO room(name, coordinate_x, coordinate_y, floor, disponibility) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(salle.Nom, salle.Coordonnees_x, salle.Coordonnees_y, salle.Etage, salle.Disponibilite, salle.Photo)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteSalle(id int) {
	stmt, err := config.Db().Prepare("DELETE FROM room WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateSalle(salle Salle) {
	stmt, err := config.Db().Prepare("UPDATE room SET name = ?, coordinate_x = ?, coordinate_y = ?, floor = ?, disponibility = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(salle.Nom, salle.Coordonnees_x, salle.Coordonnees_y, salle.Etage, salle.Disponibilite, salle.Id)
	if err != nil {
		log.Fatal(err)
	}
}
