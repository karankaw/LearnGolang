package main

import "fmt"

func main() {
	fmt.Printf("Switch Example switch statement;expression\n")

	Main()

}

func Main() {
	//var cities = [10]string {} // One Way
	//cities := [...]string {"Pune", "Navi Mumbai", "Delhi", "London", "Barcelona", "Paris", "Dubai"	}
	var city string
	var ptrCity *string
	ptrCity = &city
	fmt.Print("Enter City\n")
	fmt.Scan(ptrCity)
	fmt.Println()

	switch abrakadabra := "ABRAKADABRA"; city {
	case "Delhi":
		fmt.Println("Delhi is capital of India")
		fallthrough
	case "Mumbai":
		fmt.Println("Mumbai is financial hub of India")
		fmt.Println()
	default:
		fmt.Printf("Other City has no match %s\n", abrakadabra)
		break
	}

	//No Expression means boolean condition in case labels
	//Only Statement means statement;expression
	//switch statement; {
	//switch expression {
	switch nationalCaptial := "Delhi"; {
	case city == nationalCaptial:
		fmt.Println("Delhi Found!")
		fallthrough
	default:
		fmt.Println("Found Nothing")
	}

	var wholeNum int
	fmt.Println("Please enter wholeNum")
	fmt.Scan(&wholeNum)
	switch wholeNum {
	case 23:
		fmt.Println("23 Entered!")
	default:
		fmt.Println("23 NOT Entered!")
	}
 
}
