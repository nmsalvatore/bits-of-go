package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
