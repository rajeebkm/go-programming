package main

import "fmt"


func main() {
	conferenceName := "Go Conference";
	const conferenceTickets uint = 50;
	var remainingTickets uint = 50;

	// Types
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("There are a total of", conferenceTickets, "tickets and remaining", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Thank you for booking a ticket for", conferenceName)

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are remaining %v tickets available for the %v\n", remainingTickets, conferenceName)

	// Arraya $ Slices
	// var bookings_1 = []string{} // empty array
	var bookings_2 = [50]string{"Rajeeb", "Sanjay", "Pradeeo"} // Fixed Size and Fixed Type, Can update the array upto 50 elements
	//Define type of array
	// var bookings [50]string // Alternative way to declare and assign size and type of array
	// using sliced
	var bookings []string
	//Asign value
	bookings_2[0] = "Rajeeb"
	bookings_2[1] = "Sanjay"
	bookings_2[2] = "Pradeeo"

	// bookings[0] = firstName + " " + lastName
	// with slices
	bookings = append(bookings, firstName + " " + lastName)
	

	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("First bboking: %v\n", bookings[0])
	fmt.Printf("bookings type: %T\n", bookings)
	fmt.Printf("bookings size: %v\n", len(bookings))

	// Slice is an abstraction of an array (Dynamic size) variable length or get an sub-array of it's own
	// Slices are index-based and have a size, but is resized when needed


}