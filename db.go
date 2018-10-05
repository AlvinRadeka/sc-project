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

type parsedUsers struct {
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

func setVisitorCount(v int) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return err
	}
	defer conn.Close()

	// call producer here
	// move redis SET from here to nsq consumer
	_, err = conn.Do("SET", "visitor_count", v)
	if err != nil {
		return err
	}

	return nil
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

func handleGet() ([]parsedUsers, int) {
	visitorCount, err := getVisitorCount()
	if err != nil {
		log.Fatalln(err)
	}

	visitorCount = visitorCount + 1
	err = setVisitorCount(visitorCount)
	if err != nil {
		log.Fatalln(err)
	}

	rawUsers := getUsers()
	users := []parsedUsers{}
	t := time.Now().Year()

	for _, v := range rawUsers {
		user := parsedUsers{}
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

	return users, visitorCount
}

func getUsers() []*users {
	fmt.Println("# Started Reading Users")

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	users := []*users{}
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
		log.Fatalln(err)
	}

	fmt.Println("# Finished Reading Users")

	return users
}

func getFilteredUsers(filter string) []*users {
	fmt.Println("# Started Reading Filtered Users")

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

	fmt.Println("# Finished Reading Filtered Users")

	return users
}
