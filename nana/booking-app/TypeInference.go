package main

import (
	"fmt"
)
//Either declare or initialize
func autoInference() {
	var dontInitDontDeclare = ":-)"
	var username string            //Type Provided here
	var event = "ToBeDecided"      //Type Inference
	const numTickets = -2.01 //int uint int32 int64 float32 bool
	var remainingTickets uint
	//remainingTickets = -20   //Assigning signed to this also raises error

	username = "KK"
	event = "Pub"
	fmt.Printf("Types %v %T  %v \n",numTickets, numTickets, dontInitDontDeclare)
	fmt.Printf("Name provided : %v booked %v tickets for %v %T\n", username, numTickets, event, event)
	fmt.Println(remainingTickets)
}

func AutoInference() {
	//Same Name but Capital Initial Letter, still allowed Public Method
	fmt.Println("Public AutoInference called, not autoinference private")
}
