package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var fullName string
	var email string
	var userTickets int
	// ask users for their name
	fmt.Println("Enter you name: ")
	fmt.Scan(&fullName)

	fmt.Println("Enter you email: ")
	fmt.Scan(&email)

	fmt.Println("Enter you number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", fullName, userTickets, email)
}
