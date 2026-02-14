package main

import (
	"errors"
	"fmt"
	"os"
)

type InappropriateGreetingError struct {
	greeting string
}

func (e *InappropriateGreetingError) Error() string {
	return fmt.Sprintf("inappropriate greeting %q", e.greeting)
}

func respondTo(greeting string) (string, error) {
	if greeting == "you're dumb" {
		return "", &InappropriateGreetingError{greeting}
	}
	return "Hello, friend!", nil
}

func loadResponse() (string, error) {
	response, err := respondTo("you're dumb")
	if err != nil {
		return "", fmt.Errorf("load response: %w", err)
	}
	return response, nil
}

func main() {
	response, err := loadResponse()

	var ige *InappropriateGreetingError
	if err != nil && errors.As(err, &ige) {
		fmt.Fprintf(os.Stdout, "dickhead detected: %v\n", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(response)
}
