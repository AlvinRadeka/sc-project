package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// PageInfo is
type PageInfo struct {
	Users        []ParsedUsers
	VisitorCount int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	users, visitorCount := HandleGet()

	PageInfo := PageInfo{
		users,
		visitorCount,
	}

	template, err := template.ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, PageInfo)
}

//
func filterHandler(w http.ResponseWriter, r *http.Request) {
	filteredUsers := GetFilteredUsers("alvin")
	fmt.Println(filteredUsers)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.HandlePost()
	//database.HandleGet()
}
