package main

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "root"
// dbname   = "postgres"

type users struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Msisdn      string    `db:"msisdn"`
	Email       string    `db:"email"`
	BirthDate   time.Time `db:"birth_date"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
}

// ParsedUsers is for converting nil or other value to recognizable form
type ParsedUsers struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Msisdn      string `json:"msisdn"`
	Email       string `json:"email"`
	BirthDate   string `json:"birth_date"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
	Age         int    `json:"age"`
	Calculation string `json:"calculation"`
}

func getVisitorCount() (int, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	value, err := redis.Int(conn.Do("GET", "visitor_count"))
	if err != nil {
		return 0, err
	}

	return value, nil
}

func convertUsers(rawUsers []*users) []ParsedUsers {
	users := []ParsedUsers{}
	t := time.Now().Year()

	for _, v := range rawUsers {
		user := ParsedUsers{}
		user.ID = v.ID
		if v.Name == "" {
			user.Name = "-"
		} else {
			user.Name = v.Name
		}

		if v.Msisdn == "" {
			user.Msisdn = "-"
		} else {
			user.Msisdn = v.Msisdn
		}

		if v.Email == "" {
			user.Email = "-"
		} else {
			user.Email = v.Email
		}

		if v.BirthDate.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.BirthDate = v.BirthDate.Format("2006-01-02")
			user.Age = t - v.BirthDate.Year()
		} else {
			user.BirthDate = "-"
			user.Age = 0
		}

		if v.CreatedTime.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.CreatedTime = v.CreatedTime.Format("2006-01-02 15:04:05")
		} else {
			user.CreatedTime = "-"
		}

		if v.UpdatedTime.Format("2006-01-02 15:04:05") != "0001-01-01 00:00:00" {
			user.UpdatedTime = v.UpdatedTime.Format("2006-01-02 15:04:05")
		} else {
			user.UpdatedTime = "-"
		}

		user.Calculation = "-"

		users = append(users, user)
	}

	return users
}

func handleGet() ([]ParsedUsers, int) {
	visitorCount, err := getVisitorCount()
	if err != nil {
		log.Fatalln(err)
	}

	visitorCount = visitorCount + 1
	err = producer(visitorCount)
	if err != nil {
		log.Fatalln(err)
	}

	rawUsers := getUsers("")
	users := convertUsers(rawUsers)

	return users, visitorCount
}

func handleFilter(filter string) []ParsedUsers {
	rawUsers := getUsers(filter)
	users := convertUsers(rawUsers)

	// sentUsers, err := json.Marshal(users)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	return users
}

func getUsers(filter string) []*users {
	fmt.Println("# Started Reading Users")

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	users := []*users{}
	filter = "%" + filter + "%"
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
	WHERE name like $1
	ORDER BY id
	LIMIT 20
	`

	err = db.Select(&users, queryString, filter)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("# Finished Reading Users")

	return users
}
