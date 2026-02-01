package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("rendering template: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
