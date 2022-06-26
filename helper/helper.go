// When we have multiple packages, we should actually create folders
// for them and then put all the files belonging to that package
// in that folder

package helper

import "strings"

// Capitalize first letter of func name
// to EXPORT the func :v
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}