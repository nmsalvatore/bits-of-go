package main

import (
	"fmt"
	"log"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "hello123",
	})
	fmt.Fprintf(w, "cookie set")
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Fprint(w, "no cookie found")
		return
	}
	fmt.Fprintf(w, "cookie value: %s", cookie.Value)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/read", ReadCookie)
	mux.HandleFunc("/set", SetCookie)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
