package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter hte number of tickets you want to order: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v %v has booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("Your tickets will be sent to your email address: %v", email)
}
