package hello

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("unexpected encode error: %v", err)
	}
}
