package pkg

//import "net"
import "fmt"

func demoFunction1() {
	fmt.Println("Private : Unused Import Causes Error, Seriously ?")
}

func DemoFunction() {
	fmt.Println("Public : Unused Import Causes Error, Seriously ?")
}
