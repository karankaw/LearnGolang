package main

import "fmt"

func bookLogic() {
	const totalTickets int = 50
	var userTickets int
	var remainingTickets uint = 50
	var bookings [50]string //{"Nana", "Nicole", "Peter"}

	bookings[0] = "Nana"
	bookings[1] = "Nicole"

	fmt.Println("Total No. of tickets", totalTickets)

	var firstName string
	var lastName string
	fNamePtr := &firstName


	var rates = [...]int {12,33}
	fmt.Println(rates[0])

	fmt.Printf("Enter FirstName : ")
	fmt.Scan(fNamePtr)

	fmt.Printf("\nEnter LastName : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Number of Tickets You want to buy : ")
	fmt.Scan(&userTickets)

	bookings[2] = firstName + " " + lastName

	fmt.Printf("Thanks %v %v for booking %v tickets   ", firstName, lastName, userTickets)


	remainingTickets = remainingTickets - uint(userTickets) //Casting Int to uint
	fmt.Printf("\nTickets Remaining : %v\n", remainingTickets)
}
