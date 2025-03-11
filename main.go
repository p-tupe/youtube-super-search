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

var templates = template.Must(template.ParseFiles(
	"./public/index.html",
	"./public/results.html",
))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			templates.Execute(w, "index.html")
		} else if r.Method == http.MethodPost {
			items, err := queryYoutube(r.FormValue("q"))
			if err != nil {
				fmt.Println(err)
				// TODO: expand on error page
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = templates.ExecuteTemplate(w, "results.html", items)
			if err != nil {
				fmt.Println("Error : ", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	fmt.Println(time.Now(), " Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
