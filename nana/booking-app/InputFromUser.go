package main

import (
	"fmt"
)

func scanInput(){
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)
	fmt.Println(name)
}