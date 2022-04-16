package module1

import "fmt"

func AssignArrays(arr5 [5]int, arr2 [2]int) bool {
	fmt.Printf("%v\n", arr5)
	// arr5 = arr2
	return true
}

func Add(a, b int) int {
	return a + b
}
