package main

//import "net/http"

// "fmt"
import "sample.com/module1"
// "github.com/karankaw/booking-app/pkg"
// "github.com/karankaw/booking-app/pkg/nested"

func main() {
	var arrSize5 [5]int
	arrSize2 := [2]int{}

	module1.AssignArrays(arrSize5,arrSize2)
}

//func main1(){
// func main() {
//print("Hello World!")

// pkg.DemoFunction()
// nested.Display()

//InterPolation()

//Demo()

//formatter()

// autoInference()

// AutoInference()

// showShortAssignmentForAutoInferenceVar()

// scanInput()

// showPointers()

// bookLogic()

// fmt.Println("Sum is :", arraysDemo.Add(1, 2))

//}

/*
func InterPolation() {
	var confName = "Go Conference"
	fmt.Println("Welcome to", confName, "booking application!")
}
*/

/*
The space automatically gets added after variable Name

go list

go run <Name>.gog

No Semicolon

no commas in list - import (""
                            )

package pkg
       package nested

import modulename/package/package

package.FUNCNamePublic

Public Name - Capital Case

go run File1.go File2.go

Only 1 main func per package

go run github.com/karankaw/booking-app


*/
//go run hello.go
