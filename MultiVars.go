package main

import "fmt"

func main() {
	//Multiple Var Declaration, Same Types

	var i, j, k int16 = 65, 66, 67                //Same Type , You can mention Explicitly for Multiple Variables
	var x, y, z = 1.05, "Hulo-Mundo", '$'         //Type Inference, Multiple Variables
	p, q, r := 7.392, "Foo--BAR", '%'             //Shorthand, Diff Type, Type Inference
	colr1, colr2, colr3 := "RED", "GREEN", "BLUE" //Shorthand, Same Type, Type Inference

	fmt.Println(x, y, z)
	fmt.Println(i, j, k)
	fmt.Println(p, q, r)
	fmt.Println(colr1, colr2, colr3)
	fmt.Printf("%c %c %c\n", i, j, k)
	fmt.Printf("%d %d %d\n", i, j, k)

	//MultiVar Declaration Different Types - Explicit without Initialisation

	var (
		num8  int8
		num32 float32
		numUn uint16
	)
	num8 = -128 // 127 and -128
	num32 = 58988.873450643
	numUn = 9680

	fmt.Println(num8, num32, numUn)
	fmt.Printf("%T %T %T\n", num8, num32, numUn)

	//MultiVar Declaration Different Types - Explicit with Initialisation
	var (
		realNum8  float32 = 8.089
		realNum64 float64 = 897646464664646.00976554431
		char      byte    = 'Ö'
	)
	fmt.Println(realNum8, realNum64, char)

}
