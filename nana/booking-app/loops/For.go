package main

import (
	"fmt"
)

var x = 509 //Unused var does not happen outside func
//  y := 509   //Shorthand assignment does not work outside func

func main() {
	fmt.Println("Looping")

	fruits := [10]string{"Banana", "Orange", "Kiwi", "Apple"}
	fruits[4] = "Papaya"
	fruits[5] = "WaterMelon"
	fruits[7] = "Cherry"

	fmt.Printf("Length : %d\n", len(fruits))

	for i, str := range fruits {
		fmt.Printf("%s is at index %d\n", str, i)
	}

	for _, str := range fruits {
		fmt.Printf("Fruit %s", str)
	}

	for i := range fruits{
		fmt.Printf("%d Index has value %s",i,fruits[i])
	}

	mix := [5]float32{} //You need to provide type and size to Array
	var hetero = [10]rune{}
	mix[2] = float32(3.209)
hetero[1] = 'B' //Space Agnostic , Tab agnostic
	fmt.Println(mix, hetero)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(numbers); i++ {
		if i > 3 {
			fmt.Println(i, "is greater than 3")
		} else { //else always starts at same line as preceding closing brace
			fmt.Println(i, "is less than or equals 3")
		}
	}

	j := 0
	for j < len(numbers) { //Simulating While loop, go has no "while" construct
		if j > 3 {
			fmt.Println(j, "is greater than 3")
		} else { //else always starts at same line as preceding closing brace
			fmt.Println(j, "is less than or equals 3")
		}
		j++
	}

}
