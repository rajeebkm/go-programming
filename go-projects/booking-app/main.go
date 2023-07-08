package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference";
	const conferenceTickets uint = 50;
	var remainingTickets uint = 50;
	bookings := []string{} // slice

	// Types
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("There are a total of", conferenceTickets, "tickets and remaining", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Thank you for booking a ticket for", conferenceName)

	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint
	// // ask user for their name
	// fmt.Println("Enter your first name:")
	// fmt.Scan(&firstName)
	// fmt.Println("Enter your last name:")
	// fmt.Scan(&lastName)
	// fmt.Println("Enter your email:")
	// fmt.Scan(&email)
	// fmt.Println("Enter your number of tickets to book:")
	// fmt.Scan(&userTickets)

	// remainingTickets = remainingTickets - userTickets

	// fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// fmt.Printf("There are remaining %v tickets available for the %v\n", remainingTickets, conferenceName)

	// // Arraya $ Slices
	// // var bookings_1 = []string{} // empty array
	// var bookings_2 = [50]string{"Rajeeb", "Sanjay", "Pradeeo"} // Fixed Size and Fixed Type, Can update the array upto 50 elements
	// //Define type of array
	// // var bookings [50]string // Alternative way to declare and assign size and type of array
	// // using sliced
	// var bookings []string
	// // or,
	// // var bookings = []string{}
	// //or, 
	// // bookings := []string{}

	// //Asign value
	// bookings_2[0] = "Rajeeb"
	// bookings_2[1] = "Sanjay"
	// bookings_2[2] = "Pradeeo"

	// // bookings[0] = firstName + " " + lastName
	// // with slices
	// // Slice is an abstraction of an array (Dynamic size) variable length or get an sub-array of it's own
	// // Slices are index-based and have a size, but is resized when needed
	// bookings = append(bookings, firstName + " " + lastName)


	// fmt.Printf("List of bookings: %v\n", bookings)
	// fmt.Printf("First bboking: %v\n", bookings[0])
	// fmt.Printf("bookings type: %T\n", bookings)
	// fmt.Printf("bookings size: %v\n", len(bookings))

	// fmt.Printf("These are oll our bookings: %v\n", bookings)

	//Loops
	// for loops
	// 1. infinite loop
	for {
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
		// Handling errors
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// Arraya $ Slices
			// var bookings_1 = []string{} // empty array
			// var bookings_2 = [50]string{"Rajeeb", "Sanjay", "Pradeeo"} // Fixed Size and Fixed Type, Can update the array upto 50 elements
			//Define type of array
			// var bookings [50]string // Alternative way to declare and assign size and type of array
			// using sliced
			// var bookings []string
			// or,
			// var bookings = []string{}
			//or, 
			// bookings := []string{}
			

			//Asign value
			// bookings_2[0] = "Rajeeb"
			// bookings_2[1] = "Sanjay"
			// bookings_2[2] = "Pradeeo"

			// bookings[0] = firstName + " " + lastName
			// with slices
			// Slice is an abstraction of an array (Dynamic size) variable length or get an sub-array of it's own
			// Slices are index-based and have a size, but is resized when needed
			bookings = append(bookings, firstName + " " + lastName)
			fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are remaining %v tickets available for the %v\n", remainingTickets, conferenceName)
			fmt.Printf("These are oll our bookings: %v\n", bookings)


			// fmt.Printf("List of bookings: %v\n", bookings)
			// fmt.Printf("First booking: %v\n", bookings[0])
			// fmt.Printf("bookings type: %T\n", bookings)
			// fmt.Printf("bookings size: %v\n", len(bookings))

			// for-each loop
			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking) // returns a slice
				fmt.Printf("The names slice is: %v\n", names)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// var noTicketsremaining bool = remainingTickets == 0
			// noTicketsremaining := remainingTickets == 0
			// If-Else
			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out, Come back next year.\n")
				break
			}
		
		} else {
			fmt.Printf("Sorry, We have only %v tickets remaining. You can't book %v tickets.\n", remainingTickets, userTickets)
			// break
			// continue
		}
	
	}
}