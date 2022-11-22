package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var isRunning bool = true

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var bookings []string

	for isRunning {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Please enter the number of tickets you want to order: ")
		fmt.Scan(&userTickets)

		remainingTickets -= uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v %v has booked %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("Your tickets will be sent to your email address: %v\n", email)
		fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Printf("Bookings: %v\nNumber of Bookings: %v\n", bookings, len(bookings))

		var willContinue int

		fmt.Print("Do you want to continue booking?\n1.Yes\n2.No\n")
		fmt.Scan(&willContinue)

		if willContinue == 2 {
			isRunning = false
		}
	}

}
