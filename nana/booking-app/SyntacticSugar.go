package main

import(
	"fmt"
)

func showShortAssignmentForAutoInferenceVar(){
	// var x 
	var p int  //Declare and Init Later
	var q float64 = 1789.056 //Declare and Initialize
	var r = 0x23 //Only Declare , Automatic Inference
	s := true //Drop var keyword, Auto Infer, Initialize

	//ShortAssignment 
	//Only works for var
	//Assignment, Type Inference, no var declaration

	fmt.Println(p,q,r,s)
}