package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)

	fmt.Printf("listening on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
