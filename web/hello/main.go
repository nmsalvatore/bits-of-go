package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello is a handler function that writes "Hello, world!"
// to the HTTP response body.
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	// NewServeMux returns a new HTTP request multiplexer
	// which matches the URL of each incoming request against
	// the handlers registered for that multiplexer.
	mux := http.NewServeMux()

	// HandleFunc registers Hello to the multiplexer created above
	// and handles requests to "/" which writes "Hello, world!" to
	// the response.
	mux.HandleFunc("/", Hello)

	// Println prints a message that the server is listening on
	// port 8080
	fmt.Println("listening on port 8080")

	// ListenAndServe listens on port 8080 and calls Serve with
	// the mux handler where Hello is registered.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
