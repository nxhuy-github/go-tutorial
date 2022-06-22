// All our code must belong to a PACKAGE
// The first statement in Go file must be PACKAGE
package main

// import built-in package in Go
import "fmt"

// ENTRYPOINT: Go need to know where to start our program
// ENTRYPOINT is a MAIN function
// meaning that for ONE Go application, we only need ONE MAIN function bc we only have ONE ENTRYPOINT
func main() {
	// declare a variable
	var conferenceName = "Go Conference"
	// like var but cannot be changed
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Get your tickets here to attend")

	// Formatted PRINT
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n" , conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
