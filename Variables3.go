package main

// import the fmt and reflect package
import (
	"fmt"
	"reflect"
	"unsafe"
)

//var globalVariable rune
//var PublicGlobalVariable byte

func main() {
	var x uint16 = 123
	var y = 123
	z := 123
	fmt.Printf("Value of i is %v and Type of i is %T \n", x, x)
	fmt.Printf("Value of i is %v and Type of i is %T \n", y, y)
	fmt.Printf("Value of i is %v and Type of i is %T \n", z, z)

	//p = 123        //undefined p
	//p int = 123    //syntax error: unexpected int at end of statement
	//p int := 123     //unexpected int at end of statement
	//var p          //syntax error: unexpected newline, expecting type
	//var p boolean  //undefined: boolean
	//var p bool     //p declared but not used
	var p = 123
	fmt.Println(p)

	//_Variable1$ := true //invalid character U+0024 '$'

	//ch := '!@' //more than one character in rune literal
	ch := '@'
	fmt.Println("Type of ch is", reflect.TypeOf(ch))//Space is automatically added around "comma" in Println
	fmt.Printf("--- Type of ch is %T \n", ch)

	ru := '@'
	fmt.Println("Type of ru is", reflect.TypeOf(ru))
	fmt.Printf("--- Type of ru is %T \n", ru)

	var aByteChar byte = 'A' //Default is int32 (Alias = rune) , byte(unint8) is to be specified
	fmt.Printf("Using percentT', Type of 'aByteChar' : %T \n", aByteChar)
	fmt.Println("Using reflect.TypeOf(), Type of 'aByteChar' :", reflect.TypeOf(aByteChar))
	fmt.Println("Using reflect.ValueOf().Kind(), Type of 'aByteChar' :", reflect.ValueOf(aByteChar).Kind())
	fmt.Println("Size of aByteChar :", unsafe.Sizeof(aByteChar)) //Size is always listed in Number of bytes

	escapedString := "$\t$"
	unEscapedString := `$\t$`
	fmt.Println("escapedString", escapedString, "unEscapedString", unEscapedString)

	greeting := "Hello"
	greeting = "Bye"
	fmt.Println("Message is ", greeting)

	//fmt.Println("Type of ch is", reflect.TypeOf(ch)) //undefined: relect  - If we dont import reflect

}
