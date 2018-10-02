package main

import (
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	user := Users{
		id:          1,
		name:        "aku",
		msisdn:      "123",
		email:       "aku@aku.aku",
		birthDate:   t,
		createdTime: t,
		updatedTime: t,
	}
	template, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, user)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.HandlePost()
	//database.HandleGet()
}
