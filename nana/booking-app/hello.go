package main

//import "net/http"
import (
	"fmt"
	"github.com/karankaw/booking-app/pkg"
	"github.com/karankaw/booking-app/pkg/nested"
)

//func main1(){
func main() {
	//print("Hello World!")
	fmt.Print("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Welcome to Our conference booking application")

	pkg.DemoFunction()
	nested.Display()

	InterPolation()

	Demo()
}

func InterPolation() {
	var confName = "Go Conference"
	fmt.Println("Welcome to", confName, "booking application!")
}

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
