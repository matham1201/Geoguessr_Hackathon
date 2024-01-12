package models

import (
	"fmt"
	"log"

	"main.go/config"
)

type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckConnection(login string, password string) bool {
	var admin Admin
	fmt.Println("login", login)
	fmt.Println("password", password)

	rows, err := config.Db().Query("SELECT * FROM login WHERE username = ?", login)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&admin.Id, &admin.Username, &admin.Password); err != nil {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	fmt.Println("admin", admin.Id)

	if admin.Username == login && admin.Password == password {
		return true
	} else {
		return false
	}
}

func GetLengthLogin() int {
	var length int

	rows, err := config.Db().Query("SELECT COUNT(*) FROM login")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&length); err != nil {
			log.Fatal(err)
		}
	}

	return length
}
