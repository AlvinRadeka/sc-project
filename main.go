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
	// init get redis
	visitorCount := 20
	// up by 1, send to NSQ, and send to PageInfo
	// ...

	users := HandleGet()
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

func filterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sukses masuk filter")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/filter/", filterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.HandlePost()
	//database.HandleGet()
}
