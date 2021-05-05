// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"
	"goinaction/chapter5/listing71/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type User from the entities package.
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./listing71.go:15:3: cannot refer to unexported field 'email' in
	// struct literal of type entities.User

	fmt.Printf("User: %v\n", u)
}
