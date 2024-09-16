package main

import (
	"fmt"
	"log"

	"divineforge.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	// Get a greeting message and print it.
	message, err := greetings.Hello("Jane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
