package main

import "fmt"

func main() {
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	conferenceName := "GO Conference 2024"
	fmt.Printf("Welcome to %v booking app with %v tickets available!\n", conferenceName, conferenceTickets)

	var firstName string
	var lastName string

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	//fmt.Scanf("Please enter your first name: %v\n", &firstName)
	//fmt.Scanf("Please enter your last name: %v\n", &lastName)

	fmt.Printf("Thank you %v %v! Your ticket is booked!\n", firstName, lastName)
	remainingTickets--
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	
}