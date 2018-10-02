package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func HandlePost() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("# Inserting values")

	var lastInsertID int
	queryString := `
	INSERT INTO sc_project.users(
		id,
		name,
		msisdn,
		email,
		birth_date,
		created_time,
		updated_time
		) VALUES(
			$1,$2,$3,$4,$5,NOW(),NOW()) 
		returning id;
	`
	err = db.QueryRow(queryString, "1", "alvin", "54321", "alvin@gmail.com", "1993-05-19").Scan(&lastInsertID)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Println("# Last inserted id =", lastInsertID)
}

func HandleGetAll() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("# Reading values")

	queryString := `
	SELECT * FROM sc_project.users
	`

	rows, err := db.Query(queryString)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var msisdn string
		var email string
		var birthDate time.Time
		var createdTime time.Time
		var updatedTime time.Time
		err = rows.Scan(&id, &name, &msisdn, &email, &birthDate, &createdTime, &updatedTime)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		fmt.Printf("%v", id)
	}
}
