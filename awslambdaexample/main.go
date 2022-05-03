// main.go
package main

import (
	"fmt"

	"awslambdaexample/startec2"

	"github.com/aws/aws-lambda-go/lambda"
)

func startEC2() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	fmt.Println("Invoked Lambda")
	lambda.Start(startec2.StartInstances)
}

func main() {
	startEC2()
}
