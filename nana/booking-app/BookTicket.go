package main

import "fmt"

func Demo() {

	var conferenceName = "Go Conference" // Variable
	const conferenceTickets = 50         //Constant
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get Your tickets here to attend")
	fmt.Println("Total Tickets :", conferenceTickets, "Remaining Tickets :", remainingTickets)

	// conferenceTickets = 30
}
