package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

var helloTmpl = template.Must(template.ParseFiles("base.html", "hello.html"))

func Hello(w http.ResponseWriter, r *http.Request) {
	err := helloTmpl.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Fatalf("rendering template: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
