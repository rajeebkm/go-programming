package main

import "fmt"


func main() {
	var conferenceName = "Go Conference";
	const conferenceTickets = 50;
	var remainingTickets = 50;
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("There are a total of", conferenceTickets, "tickets and remaining", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Thank you for booking a ticket for", conferenceName)

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2

	fmt.Printf("user %v bough %v tickets for the Go conference. \n", userName, userTickets)

}