package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageInfo struct {
	Users        []parsedUsers
	VisitorCount int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	users, visitorCount := handleGet()

	pageInfo := pageInfo{
		users,
		visitorCount,
	}

	template, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, pageInfo)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	// filteredUsers := getFilteredUsers("a")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/filter/", filterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
