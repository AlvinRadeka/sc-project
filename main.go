package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageInfo struct {
	Users        []ParsedUsers
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
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		values := []string{}
		for _, v := range r.Form["name"] {
			values = append(values, v)
		}

		users := handleFilter(values[0])
		template, err := template.ParseFiles("page.html")
		if err != nil {
			log.Fatal(err)
		}

		visitorCount, err := getVisitorCount()
		if err != nil {
			log.Fatalln(err)
		}

		pageInfo := pageInfo{
			users,
			visitorCount,
		}

		template.Execute(w, pageInfo)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/filter", filterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
