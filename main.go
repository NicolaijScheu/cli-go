package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var isRunning bool = true

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	var bookings []string

	for isRunning {

		firstName, lastName, email, userTickets := getUserInput()

		isValidUserInput := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidUserInput {
			fmt.Println("Invalid User input.")
			continue
		}

		//book tickets
		remainingTickets, bookings = bookTickets(remainingTickets, userTickets, bookings, firstName, lastName)

		//print successful booking
		printBookingInfo(firstName, lastName, userTickets, email, conferenceTickets, remainingTickets)

		//printFirstNames
		printFirstNames(bookings)

		if remainingTickets == 0 {
			fmt.Println("The conference is sold out.")
			isRunning = false
		}
	}
}

func greetUser(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("Welcome to the %v booking system.\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here!")
}

func isNameValid(name string) bool {
	return len(name) >= 2 && !regexp.MustCompile(`\d`).MatchString(name)
}

func printFirstNames(bookings []string) {
	for index, booking := range bookings {
		firstNameSplit := strings.Split(booking, " ")
		fmt.Printf("%v. %v\n", index+1, firstNameSplit[0])
	}
}

func printBookingInfo(firstName string, lastName string, userTickets int, email string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("User %v %v has booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("Your tickets will be sent to your email address: %v\n", email)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) bool {
	firstNameIsValid := isNameValid(firstName)
	lastNameIsValid := isNameValid(lastName)

	if !firstNameIsValid || !lastNameIsValid {
		fmt.Println("Please enter a valid Name!")
	}

	emailIsValid := len(email) > 5 && strings.Contains(email, "@") && strings.Contains(email, ".")

	if !emailIsValid {
		fmt.Println("Please enter a valid Email address!")

	}

	userTicketsIsValid := userTickets > 0 && userTickets <= int(remainingTickets)

	if !userTicketsIsValid {
		fmt.Println("Invalid amount of tickets!")

	}

	if userTickets > int(remainingTickets) {
		fmt.Printf("We only have %v tickets left for sale. You can't buy %v tickets.\n", remainingTickets, userTickets)

	}
	return firstNameIsValid && lastNameIsValid && emailIsValid && userTicketsIsValid
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets int, bookings []string, firstName string, lastName string) (uint, []string) {
	remainingTickets -= uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)
	return remainingTickets, bookings
}
