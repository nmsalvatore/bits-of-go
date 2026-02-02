package main

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

var tmpl = template.Must(template.ParseFiles("login.html", "protected.html"))
var sessions = map[string]string{}

type User struct {
	Username string
	Password string
}

func generateSessionToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		user := User{
			Username: "testuser",
			Password: "password123",
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username != user.Username || password != user.Password {
			http.Error(w, "invalid username or password", http.StatusBadRequest)
			return
		}

		token := generateSessionToken()
		sessions[token] = username

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    token,
			HttpOnly: true,
		})

		http.Redirect(w, r, "/protected", http.StatusSeeOther)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid form method", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session")
	if err == nil {
		delete(sessions, cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Protected(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username, ok := sessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := map[string]string{
		"Username": username,
	}
	err = tmpl.ExecuteTemplate(w, "protected.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/protected", Protected)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
