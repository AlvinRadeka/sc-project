package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 54323
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
		returning id;`
	err = db.QueryRow(queryString, "1", "alvin", "12345", "alvin@gmail.com", "1993-05-19").Scan(&lastInsertID)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertID)

}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
