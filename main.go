package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const conferenceName = "Go Conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var isRunning bool = true
var bookings = make([]UserData, 0)

func main() {

	greetUser()

	for isRunning {

		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidUserInput := validateUserInput(firstName, lastName, email, userTickets)

		if !isValidUserInput {
			fmt.Println("Invalid User input.")
			continue
		}

		//book tickets
		bookTickets(userTickets, firstName, lastName, email)

		//print successful booking
		printBookingInfo(firstName, lastName, userTickets, email)

		//printFirstNames
		printFirstNames()

		if remainingTickets == 0 {
			fmt.Println("The conference is sold out.")
			isRunning = false
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to the %v booking system.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

func printFirstNames() {
	for key, element := range bookings {
		fmt.Printf("%v: %v, Email address: %v, Tickets: %v\n", key+1, element.firstName, element.email, element.numberOfTickets)
	}
}

func printBookingInfo(firstName string, lastName string, userTickets int, email string) {
	fmt.Printf("User %v %v has booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("Your tickets will be sent to your email address: %v\n", email)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) bool {
	firstNameIsValid := helper.IsNameValid(firstName)
	lastNameIsValid := helper.IsNameValid(lastName)

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

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= uint(userTickets)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)
}
