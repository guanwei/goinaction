// Sample program to show how to use an interface in Go.
package main

import "fmt"

// notifier is an interface that defined notification type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@example.com"}

	sendNotifiaction(u)

	// ./listing36.go:29:18: cannot use u (type user) as type
	// notifier in argument to sendNotifiaction:
	// user does not implement notifier (notify method has pointer receiver)
}

// sendNotifiaction accepts values that implement the notifier
// interface and sends notifications.
func sendNotifiaction(n notifier) {
	n.notify()
}
