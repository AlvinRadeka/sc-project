package main

import (
	"html/template"
	"log"
	"net/http"
)

// Page is
type Page struct {
	Title string
	Body  []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	users := GetUsers()
	template, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, users)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.HandlePost()
	//database.HandleGet()
}
