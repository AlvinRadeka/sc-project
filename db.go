package main

import (
	"database/sql"
	"fmt"
	"log"

	// postgres import
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// HandleGet is
func HandleGet() []Users {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	fmt.Println("# Reading values")

	queryString := `
	SELECT * FROM sc_project.users
	LIMIT 20
	`

	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	users := []Users{}
	for rows.Next() {
		user := Users{}
		err = rows.Scan(&user.ID, &user.Name, &user.Msisdn, &user.Email, &user.BirthDate, &user.CreatedTime, &user.UpdatedTime)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		users = append(users, user)
	}

	fmt.Println("# Finished Reading")

	return users
}
