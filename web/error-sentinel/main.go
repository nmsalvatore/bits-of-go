package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

var ErrNotInTheMood = errors.New("not in the mood")

func inviteToParty() error {
	res := []string{"yes", "no"}[rand.Intn(2)]
	if res == "no" {
		return fmt.Errorf("response: %w", ErrNotInTheMood)
	}
	return nil
}

func main() {
	err := inviteToParty()
	if err != nil && errors.Is(err, ErrNotInTheMood) {
		fmt.Println("Well, fine then.")
		return
	}

	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("Yay!")
}
