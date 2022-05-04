// main.go
package main

import (
	"fmt"

	"awslambdastopec2/stopec2"

	"github.com/aws/aws-lambda-go/lambda"
)

func stopEC2() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	fmt.Println("Invoked Lambda - StopEC2")
	lambda.Start(stopec2.StopInstances)
}

func main() {
	stopEC2()
}
