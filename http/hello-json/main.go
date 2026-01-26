package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Data struct {
	Message string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	data := Data{
		Message: fmt.Sprintf("Hello, %s!", name),
	}

	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
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
