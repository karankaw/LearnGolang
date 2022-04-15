package main

import (
	"fmt"
)

func main() {

	var days [7]string
	months := []string{"Jan", "Feb"}
	monthsRUNE := [12]rune{'J', 'F'}
	digits := []int{0, 1}
	floats := [...]float64{float64(2), 1.63}

	fmt.Printf("%v %v %d \n", len(months), len(digits), len(floats))

	digits[2] = 43 //Runtime Error as it computes size during execution

	// floats[2] = 54 //Compile Time Error because of [...] size is compile time computed i.e 2

	//len string, calculates Bytes not character

	fmt.Printf("%v %v %d \n", len(months), len(digits), len(floats))
	fmt.Println(monthsRUNE)
	
	fmt.Printf("Arrays Example! %v %v %v %v\n", days, months, digits, floats)

}
