// main.go
package main

import (
	"fmt"

	"awslambdastartec2/startec2"

	"github.com/aws/aws-lambda-go/lambda"
)

func startEC2() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	fmt.Println("Invoked Lambda - StartEC2")
	lambda.Start(startec2.StartInstances)
}

func main() {
	startEC2()
}
