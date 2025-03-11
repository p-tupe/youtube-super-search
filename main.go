package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Options struct {
	Title string
}

var templates = template.Must(template.ParseFiles("./public/index.html", "./public/results.html"))

func main() {
	http.HandleFunc("/", rootHandler)

	fmt.Println(time.Now(), " Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.Execute(w, "index.html")
	} else if r.Method == http.MethodPost {
		results := r.FormValue("query")
		if err := templates.ExecuteTemplate(w, "results.html", &struct{ Results string }{
			Results: "You searched for: " + results,
		}); err != nil {
			fmt.Println("Error : ", err)
		}
	}
}
