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
	// when we declare a variable and assign it a value immediately, Go can IMPLY its data-type
	// var conferenceName = "Go Conference"
	// like var but cannot be changed
	// const conferenceTickets = 50
	// var remainingTickets = 50

	// %T to print the type of variable
	// fmt.Printf("conferenceName is %T, remainingTickets is %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

	// SUGAR syntax
	conferenceName := "Go Conference"
	// but it DOESN'T if we want to create a constant
	const conferenceTickets = 50
	// and it DOESN'T too if we want to assign a TYPE for a variable
	var remainingTickets uint = 50

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Get your tickets here to attend")

	// Formatted PRINT
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n" , conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Go is statically TYPED LANGUAGE
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	// instead of passing the VALUE of userName (which is empty)
	// we pass the REFERENCE or the MEMORY ADDRESS of the var userName :))
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	// <<< Go POINTER >>>
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
	
	fmt.Println("Enter number tickets: ")
	fmt.Scan(&userTickets)

	// the reason of uint() here is bc the mismatch between remainingTickets (int) and userTickets (unint) 
	remainingTickets -= uint(userTickets)
	
	fmt.Printf("Thank you for %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the conference", remainingTickets)

}
