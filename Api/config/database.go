package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	fmt.Println("Connecting to database...")

	//conecter a MYSQL the name of the database is Geoguessr_Ynov the host is "10.44.17.117" and the port is 3306 and the user is root and the password is root
	db, err = sql.Open("mysql", "root:root@tcp(10.44.17.117:3306)/geoguessr")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database !")

}

func GetLengthRoom() int {
	var length int
	err := db.QueryRow("SELECT COUNT(*) FROM room").Scan(&length)
	if err != nil {
		log.Fatal(err)
	}
	return length
}

func GetLengthlogin() int {
	var length int
	err := db.QueryRow("SELECT COUNT(*) FROM login").Scan(&length)
	if err != nil {
		log.Fatal(err)
	}
	return length
}

func GetLengthScoreboard() int {
	var length int
	err := db.QueryRow("SELECT COUNT(*) FROM scoreboard").Scan(&length)
	if err != nil {
		log.Fatal(err)
	}
	return length
}

// Getter for db var
func Db() *sql.DB {
	return db
}
