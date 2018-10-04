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
	Age         int
	Calculation int
}

// HandleGet is
func HandleGet() []ParsedUsers {
	rawUsers := GetUsers()
	users := []ParsedUsers{}
	t := time.Now().Year()

	for _, v := range rawUsers {
		user := ParsedUsers{}
		user.ID = v.ID
		user.Name = v.Name
		user.Msisdn = v.Msisdn
		user.Email = v.Email
		if v.BirthDate.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.BirthDate = v.BirthDate.Format("2006-01-02")
			user.Age = t - v.BirthDate.Year()
		}
		if v.CreatedTime.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.CreatedTime = v.CreatedTime.Format("2006-01-02 15:04:05")
		}
		if v.UpdatedTime.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.UpdatedTime = v.UpdatedTime.Format("2006-01-02 15:04:05")
		}

		users = append(users, user)
	}

	return users
}

// GetUsers is func for get all users
func GetUsers() []*Users {
	fmt.Println("# Started Reading Users")

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	users := []*Users{}
	queryString := `
	SELECT 
		coalesce(id, '') as id,
		coalesce(name, '') as name,
		coalesce(msisdn, '') as msisdn,
		coalesce(email, '') as email,
		coalesce(birth_date, '0001-01-01 00:00:00') as birth_date,
		coalesce(created_time, '0001-01-01 00:00:00') as created_time,
		coalesce(updated_time, '0001-01-01 00:00:00') as updated_time
	FROM sc_project.users
	ORDER BY id
	LIMIT 20
	`

	err = db.Select(&users, queryString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("# Finished Reading Users")

	return users
}

// GetFilteredUsers is
func GetFilteredUsers(filter string) []*Users {
	fmt.Println("# Started Reading Users")

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	users := []*Users{}
	queryString := `
	SELECT 
		coalesce(id, '') as id,
		coalesce(name, '') as name,
		coalesce(msisdn, '') as msisdn,
		coalesce(email, '') as email,
		coalesce(birth_date, '0001-01-01 00:00:00') as birth_date,
		coalesce(created_time, '0001-01-01 00:00:00') as created_time,
		coalesce(updated_time, '0001-01-01 00:00:00') as updated_time
	FROM sc_project.users
	WHERE name like '%$1%'
	ORDER BY id
	LIMIT 20
	`

	err = db.Select(&users, queryString, filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("# Finished Reading Users")

	return users
}
