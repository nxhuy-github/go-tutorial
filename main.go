// Go programs are organized into packages
// A package is a collection of Go files

// 3 levels of Scope
// Local: declaration within function/block
// Package: declaration outside all functions -> can be used EVERYWHERE in the SAME package
// Global: declaration outside all functions & UPPERCASE first letter -> can be used EVERYWHERE ACROSS ALL packages

// The first statement in Go file must be PACKAGE
package main

// import built-in package in Go
// each package is on a new line
import (
	"fmt"
	"go-tutorial/helper"
	"sync"
	"time"
)

// Package Variable
// ERROR if use SUGAR SYNTAX: packageVariable := "package Variable" -> error
var conferenceName = "Go Conference"
// but it DOESN'T if we want to create a constant
const conferenceTickets = 50
// and it DOESN'T too if we want to assign a TYPE for a variable
var remainingTickets uint = 50

// 'type' creates a NEW type with the name you specify
// in this case: create a type called UserData based on a struct of firstName, lastName, ...
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// by default, the MAIN goroutine DOESN'T wait for other goroutines
// WaitGroup waits for the launched goroutine to finish
var wg = sync.WaitGroup{}

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
	// conferenceName := "Go Conference"
	// but it DOESN'T if we want to create a constant
	// const conferenceTickets = 50
	// and it DOESN'T too if we want to assign a TYPE for a variable
	// var remainingTickets uint = 50

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// fmt.Println("Get your tickets here to attend")
	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	// Formatted PRINT
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// ARRAY in Go is FIXED SIZE, syntax: var <var_name> = [size] var_type {}
	// for ex: var bookings  = [50] string {}
	// the SPACE here DOESN'T matter: var bookings  = [50]string{} is still valid
	// ALTERNATIVE way to declare a empty array
	// var bookings [50]string

	// SLICE in Go is DYNAMIC SIZE
	// var bookings []string
	// ALTERNATIVE way to declare a empty slice
	// var bookings = []string{}
	//bookings := []string{}

	// Create a LIST OF MAP
	//bookings := make([]map[string]string, 0)

	// Create a LIST OF UserData
	bookings := make([]UserData, 0)

	// Go only has ONE keyword for LOOP
	for {
		// Go is statically TYPED LANGUAGE
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint
		// ask user for their name
		// instead of passing the VALUE of userName (which is empty)
		// we pass the REFERENCE or the MEMORY ADDRESS of the var userName :))
		// fmt.Println("Enter your first name: ")
		// fmt.Scan(&firstName)

		// fmt.Println("Enter your last name: ")
		// fmt.Scan(&lastName)

		// fmt.Println("Enter your email: ")
		// fmt.Scan(&email)

		// <<< Go POINTER >>>
		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		// fmt.Println("Enter number tickets: ")
		// fmt.Scan(&userTickets)

		firstName, lastName , email , userTickets := getUserInput()

		// isValidName := len(firstName) >= 2 && len(lastName) >= 2 
		// var isValidEmail bool = strings.Contains(email, "@")
		// var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// the reason of uint() here is bc the mismatch between remainingTickets (int) and userTickets (unint)
			// remainingTickets -= uint(userTickets)
			//bookings[0] = firstName + " " + lastName
			// bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first element array: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// fmt.Printf("Thank you for %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			// fmt.Printf("%v tickets remaining for the conference\n", remainingTickets)
			bookTicket(&remainingTickets ,  userTickets , &bookings , firstName , lastName , email)
			// 'go' keyword starts a new goroutine which is 
			// a lightweight thread managed by the Go runtime
			
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// firstNames := []string{}
			// // 'range' iterates over elements for different data structures (NOT ONLY arrays and slices)
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	// var firstName = names[0]
			// 	firstNames = append(firstNames, names[0])
			// }
			firstNames := printFirstName(bookings)
			fmt.Printf("The first names of booking are: %v\n", firstNames)
			
			//var noTicketRemaining = remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
			
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short, please try again")
			}
			if !isValidEmail {
				fmt.Println("Your email address doesn't contain @, please try again")
			}
			if !isValidTicketNumber {
				fmt.Println("Your number tickets is INVALID, please try again")
			}
		}

	}
	wg.Wait()
	
	// city := "London"
	// switch city {
	// 	case "London":
	// 		// execute something
	// 	case "Singapore":
	// 		// execute something
	// 	case "New York":
	// 		// execute something
	// 	default:
	// 		// execute something
	// }

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint)  {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName(bookings []UserData) []string {
	firstNames := []string{}
	// 'range' iterates over elements for different data structures (NOT ONLY arrays and slices)
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
// 	var isValidEmail bool = strings.Contains(email, "@")
// 	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets *uint,  userTickets uint, bookings *[]UserData, firstName string, lastName string, email string) {
	*remainingTickets -= userTickets

	// create a MAP
	// make() helps us create an EMPTY map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	*bookings = append(*bookings, userData)
	fmt.Printf("List of bookings is %v \n", *bookings)

	fmt.Printf("Thank you for %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the conference\n", *remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string)  {
	time.Sleep(10 * time.Second) // 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("######################")
	wg.Done()
}