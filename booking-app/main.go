package main

import "fmt"

func main() {
	const CONFERENCE_TICKET = 50
	conferenceName := "Go Conference"
	var remainingTickets uint = 50
	// A slice of an array
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", CONFERENCE_TICKET, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")


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

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %v\n", bookings)
	fmt.Printf("Array Length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email %v \n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}


