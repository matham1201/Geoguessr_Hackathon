package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	//conecter a MYSQL

	db, err = sql.Open("mysql", "rot@10.44.17.117(localhost:3306)/Geoguessr_Ynov_BDD")

	if err != nil {
		log.Fatal(err)
	}

}

// Getter for db var
func Db() *sql.DB {
	return db
}
