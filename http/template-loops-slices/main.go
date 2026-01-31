package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	data := []string{"apples", "oranges", "bananas"}

	err = tmpl.Execute(w, data)
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
