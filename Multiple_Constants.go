//Multiple Constants
//Multiple Individual Imports
package main

import "fmt"
import "reflect"
	
const (
	Http_Ok, Http_Moved_Permanently, Http_Unauthorized = "200", "304", 401
	Http_Internal_Server_Error                         = "500"
	Pi                                                 = 3.14
)

func main() {
	fmt.Println("Multiple Constants")
	fmt.Println(Http_Ok, reflect.TypeOf(Http_Ok))
	fmt.Println(Http_Moved_Permanently, reflect.TypeOf(Http_Moved_Permanently))
	fmt.Println(Http_Unauthorized, reflect.TypeOf(Http_Unauthorized))
	fmt.Println(Http_Internal_Server_Error, reflect.TypeOf(Http_Internal_Server_Error))
	fmt.Println(Pi, reflect.TypeOf(Pi))
}
