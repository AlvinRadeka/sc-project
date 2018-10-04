package main

import (
	"fmt"
	"log"
	"time"

	// postgres import
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "root"
// dbname   = "postgres"

// Users is
type Users struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Msisdn      string    `db:"msisdn"`
	Email       string    `db:"email"`
	BirthDate   time.Time `db:"birth_date"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
}

// ParsedUsers is
type ParsedUsers struct {
	ID          int
	Name        string
	Msisdn      string
	Email       string
	BirthDate   string
	CreatedTime string
	UpdatedTime string
}

// HandleGet is
func HandleGet() []ParsedUsers {
	users := []ParsedUsers{}

	return users
}

// GetUsers is func for get all users
func GetUsers() []Users {
	fmt.Println("# Started Reading Users")

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	users := []Users{}
	queryString := `
	SELECT * FROM sc_project.users
	LIMIT 20
	`

	err = db.Select(&users, queryString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("# Finished Reading Users")

	return users
}
