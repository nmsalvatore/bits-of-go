package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

type App struct {
	tmpl *template.Template
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	a.tmpl.ExecuteTemplate(w, "home.html", nil)
}

func (a *App) About(w http.ResponseWriter, r *http.Request) {
	a.tmpl.ExecuteTemplate(w, "about.html", nil)
}

func (a *App) Contact(w http.ResponseWriter, r *http.Request) {
	a.tmpl.ExecuteTemplate(w, "contact.html", nil)
}

func main() {
	mux := http.NewServeMux()

	app := App{
		tmpl: template.Must(template.ParseFiles("about.html", "contact.html", "home.html")),
	}

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/contact", app.Contact)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
