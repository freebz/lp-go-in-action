// Sample program to show how to use an interface in Go.
package main

import (
    "fmt"
)

// notifier is an interface that defined notification
// type behavior.
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
    u := user{"Bill", "bill@email.com"}
    
    sendNotification(&u)
    u.notify()
    (&u).notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notificiations.
func sendNotification(n notifier) {
    n.notify()
}
