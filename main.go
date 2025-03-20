package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates
var templateFS embed.FS

var templates = template.Must(template.ParseFS(templateFS, "templates/*.html"))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			templates.Execute(w, "index.html")
		} else if r.Method == http.MethodPost {
			items, err := queryYoutube(r.FormValue("q"))
			if err != nil {
				fmt.Println(err)
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

	fmt.Println(time.Now(), " Server listening on port 8084...")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
