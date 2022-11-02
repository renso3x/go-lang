package main

import "fmt"

func main() {
	const CONFERENCE_TICKET = 50
	conferenceName := "Go Conference"
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", CONFERENCE_TICKET, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")

	
	// declare variable with data-type
	var userName string
	var userTickets int


	// Get user value by -> using pointer `&variable`
	fmt.Scan(&userName)

	fmt.Printf("User %v booked %v tickets. \b", userName, userTickets)

}


