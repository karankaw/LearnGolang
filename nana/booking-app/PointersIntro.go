package main

import (
	"fmt"
)

func showPointers() {
	ticketsNum := 20
	var pointer *int = &ticketsNum

	country := "INDIA"
	var strPointer *string
	strPointer = &country

	pincode := "411021" //'' is rune datatype in golang
	strPointer = &pincode
	newPointerShrtHand := &pincode

	fmt.Printf("Pointer Type : %T\n", pointer)
	fmt.Printf("Pointer Value  : %v\n", pointer)
	fmt.Printf("Value pointed to by Pointer  : %v\n", *pointer)

	fmt.Printf("Pointer Type : %T\n", strPointer)
	fmt.Printf("Pointer Value  : %v\n", strPointer)
	fmt.Printf("Value pointed to by Pointer  : %v\n", *strPointer)

	fmt.Printf("Value pointed to by Pointer  : %v\n", *newPointerShrtHand)

}
