package main

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

const port = "8080"

var templates = template.Must(template.ParseFiles("login.html", "protected.html"))
var sessions = map[string]string{}

type User struct {
	Username string
	Password string
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cookie, err := r.Cookie("session")
		if err == nil {
			_, ok := sessions[cookie.Value]
			if ok {
				http.Redirect(w, r, "/protected", http.StatusSeeOther)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:   "session",
				Value:  "",
				MaxAge: -1,
			})
		}

		token := generateToken()
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    token,
			HttpOnly: true,
		})

		data := map[string]string{
			"CSRFToken": token,
		}

		err = templates.ExecuteTemplate(w, "login.html", data)
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

		formToken := r.PostFormValue("csrf_token")
		cookie, err := r.Cookie("csrf_token")
		if err != nil || formToken != cookie.Value {
			http.Error(w, "invalid csrf token", http.StatusForbidden)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "csrf_token",
			Value:  "",
			MaxAge: -1,
		})

		user := User{
			Username: "testuser",
			Password: "password123",
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if username != user.Username || password != user.Password {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		token := generateToken()
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

	err = templates.ExecuteTemplate(w, "protected.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/protected", Protected)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
