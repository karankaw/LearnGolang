//Constants cannot be declared using the := syntax as its a shorthand for var and Type Inference/Untyped type

package main

import (
	"fmt"
)

func main() {
	const pi1 float32 = 3.14
	const pi2 = 3.14

	//const pi3 := 3.14 //syntax error: unexpected :=, expecting =

	pi := 3.14

	fmt.Println(pi1, pi2, pi)
	fmt.Printf("%T,%T,%T", pi1, pi2, pi)
}
