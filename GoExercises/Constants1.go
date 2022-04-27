package main

import "fmt"

const Global_PI = 3.14

func main() {

	const pi float32 = 3.1412
	const dollar = '$'

	fmt.Printf("Type of %v is %T \n", pi, pi)
	fmt.Printf("Type of %v is %T \n", dollar, dollar)
	fmt.Println(Global_PI)

	//dollar = '#' //Error, cannot assign to dollar (declared const)

	anotherF()
}

func anotherF() {
	//Global_PI = 4.0 //cannot assign to Global_PI (declared const)
	fmt.Println("Value of Global_PI :", Global_PI)
}
