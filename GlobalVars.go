
package main

import "fmt"

//global but not public
//foo := "Hello World"   //syntax error: non-declaration statement outside function body
//foo = "Hello World"    //syntax error: non-declaration statement outside function body
var foo = "Hello World"
var foo1 string = "Hulo Mundo"

var bar string = "Alright" //No error for global variables if unused
//bar = "Bye, Bye"        //Cannot Assign Values to global variable as a statement if declared before
var bye = "Bye, Bye"
var sayonara = "Sayonara"

//global and public
//Foo = "Hello World"

func main() {
	//overrides global scoped variables of same name
	//var bar1 string //bar1 declared but not used

	display()
}

func display() {

	fmt.Println(foo, foo1)
	fmt.Println(bar)
	fmt.Println(bye)

}
