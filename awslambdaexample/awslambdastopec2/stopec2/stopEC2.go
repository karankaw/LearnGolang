// main.go
package stopec2

import (
	"fmt"
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)



func StopInstances(_ context.Context, event events.CloudWatchEvent) {

	fmt.Printf("SOURCE: %+v\n", event.Source)
	fmt.Printf("ACTION: %+v\n", event.DetailType)

	const DRYRUN_OPERATION = "DryRunOperation"

	fmt.Println("Stopping EC2 Instances at EOD")

	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String("i-065ddea1dbba7a916"),
			aws.String("i-05f17330cb77ca0b9"),
		},
		DryRun: aws.Bool(true),
	}

	result, err := ec2Svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)

	// DryRun OK, Run Actual
	if ok && awsErr.Code() == DRYRUN_OPERATION {
		input.DryRun = aws.Bool(false)
		fmt.Println("Stopping")
		result, err = ec2Svc.StopInstances(input)

		if err != nil {
			fmt.Println("Error", err)

		} else {
			fmt.Println("Success", result.StoppingInstances)
		}
	}

}
