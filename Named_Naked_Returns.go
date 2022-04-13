package main

import "fmt"

func split(sum int) (x int, y float32 ) {
	//x and y will be looked up inside for automatic/naked return/named return
	x = sum * 4 / 9
	z := 265
	y = float32(sum - x)
	
	fmt.Println(z)
	return  // it automatically returns whatever named params are in return declaration
}

func main() {
	fmt.Println(split(17))
}

