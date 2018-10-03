package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Page is
type Page struct {
	Title string
	Body  []byte
}

// Users is
type Users struct {
	ID          int
	Name        string
	Msisdn      string
	Email       string
	BirthDate   time.Time
	CreatedTime time.Time
	UpdatedTime time.Time
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	users := HandleGet()
	template, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	template.Execute(w, users)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.HandlePost()
	//database.HandleGet()
}
