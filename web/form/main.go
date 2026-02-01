package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

var formTmpl = template.Must(template.ParseFiles("form.html"))

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := formTmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		message := r.PostFormValue("message")
		if message == "" {
			http.Error(w, "message can't be empty", http.StatusBadRequest)
			return
		}

		fmt.Println(message)
		http.Redirect(w, r, "/form", http.StatusSeeOther)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", Form)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
