package main

import (
	"fmt"
	"strings"
)

// Package level variables
const conferenceTickets uint = 50
var conferenceName string = "Go Conference"
var remainingTickets uint = 50;
var bookings = []string{} // slice

func main() {
	
	greetUser();

	//Loops
	// for loops
	// 1. infinite loop
	for {

		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lastName, email, userTickets)
		
		if isValidName && isValidEmail && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames();
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// If-Else
			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out, Come back next year.\n")
				break
			}
		
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain an @ sign.")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			// break
			// continue
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("There are a total of", conferenceTickets, "tickets and remaining", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
		// for-each loop
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking) // returns a slice
			fmt.Printf("The names slice is: %v\n", names)
			firstNames = append(firstNames, names[0])
		}

		return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets to book:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}


func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are remaining %v tickets available for the %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all of our bookings: %v\n", bookings)
}