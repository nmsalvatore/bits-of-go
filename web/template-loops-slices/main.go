package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

var tmpl = template.Must(template.ParseFiles("home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	data := []string{"apples", "oranges", "bananas"}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("rendering template: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
