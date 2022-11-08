package main

import (
	"fmt"
	"strings"
)

const CONFERENCE_TICKET = 50
var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUsers()

	// Infite loop | while loop
	for remainingTickets > 0 && len(bookings) < 50 {
		// Get User Input
		firstName, lastName, email, userTickets := getUserInput()
		// User Validation
		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)
		
		if !isValidName || !isValidEmail ||! isValidTickets {
			if !isValidName {
				fmt.Println("Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is not valid")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
		} else {
			if userTickets <= remainingTickets {
				makeBooking(userTickets, firstName, lastName, email)	

				firstNames := printFirstNames()

				fmt.Printf("These are the first names of the bookings: %v \n", firstNames)

				noTicketsRemaining := remainingTickets == 0
				if noTicketsRemaining {
					// end the program
					fmt.Println("Our conference is sold out. Come back next year")
					break
				}
			} else {
				fmt.Printf("We only have %v tickets remaining, you can't book %v tickets\n", remainingTickets, userTickets)
				fmt.Printf("Please try again. \n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", CONFERENCE_TICKET, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
}

func printFirstNames() []string {
	// Loop slices or arrays
	firstNames := []string{}
	// for loop
	for _, booking := range bookings {
		var names = strings.Fields(booking) // splits white space
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}


func getUserInput() (string, string, string, uint){
	// declare variable with data-type
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Get user value by -> using pointer `&variable`
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func makeBooking(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email %v \n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, CONFERENCE_TICKET)
}