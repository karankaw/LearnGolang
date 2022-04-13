package main

import "fmt"

func main() {
	fmt.Println("Main Called!")

	//inner_func1() // Does not run here, before declaration
	inner_func1 := func() {
		fmt.Println("Inner Function1 Called!")
	}

	var inner_func2 = func() {
		fmt.Println("Inner Function2 Called!")
	}

	inner_func1()
	inner_func2()
}
