package main

import (
	"fmt"

	"github.com/Akila-1998/golang/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
