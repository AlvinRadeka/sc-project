package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

// Users is
type Users struct {
	id          int
	name        string
	msisdn      string
	email       string
	birthDate   time.Time
	createdTime time.Time
	updatedTime time.Time
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// HandlePost is
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

// HandleGet is
func HandleGet() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
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
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		users := Users{}
		err = rows.Scan(&users.id, &users.name, &users.msisdn, &users.email, &users.birthDate, &users.createdTime, &users.updatedTime)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		fmt.Printf("%v \n", users)
	}
}
