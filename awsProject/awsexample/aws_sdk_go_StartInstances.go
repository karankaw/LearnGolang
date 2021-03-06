package awsexample

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// References
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ec2-example-manage-instances.html
// https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/example_code/ec2/describing_instances.go
// https://github.com/aws/aws-sdk-go/blob/v1.44.1/service/ec2/examples_test.go

// https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
// https://yourbasic.org/golang/structs-explained/
// https://gobyexample.com/structs



func StartEC2() {
	fmt.Println("Start EC2 Instances")
	getAWSSessionStart()
}

func getAWSSessionStart() {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	instance_state_name := aws.String("instance-state-name")

	values := []*string{aws.String("running")}

	ec2_Filter := &ec2.Filter{Name: instance_state_name, Values: values}

	input := new(ec2.DescribeInstancesInput)

	input.Filters = []*ec2.Filter{ec2_Filter}

	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success")
	}

	for index, reservation := range result.Reservations {
		fmt.Println("-------------------- Index : ", index, " --------------------")
		// fmt.Println(reservation.Instances)

		for _, instance := range reservation.Instances {
			fmt.Println(*instance.InstanceId)
		}
	}

	//    * instance-state-code - The state of the instance, as a 16-bit unsigned
	//    integer. The high byte is used for internal purposes and should be ignored.
	//    The low byte is set based on the state represented. The valid values are:
	//    0 (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping),
	//    and 80 (stopped).
	//
	//    * instance-state-name - The state of the instance (pending | running |
	//    shutting-down | terminated | stopping | stopped).

	if os.Args[1] == "START" { //Test i-065ddea1dbba7a916
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{ aws.String("i-065ddea1dbba7a916"),aws.String("i-05f17330cb77ca0b9")},
			DryRun: aws.Bool(true),
		}
		result, err := ec2Svc.StartInstances(input)
		awsErr, ok := err.(awserr.Error)

		fmt.Println(result)
		fmt.Println(err)
		fmt.Println(awsErr)
		fmt.Println(ok)

		fmt.Println(awsErr.Code())

		const DRYRUN_OPERATION = "DryRunOperation"

		if ok && awsErr.Code() == DRYRUN_OPERATION {
			input.DryRun = aws.Bool(false)
			fmt.Println("Starting")
			result, err = ec2Svc.StartInstances(input)

			if err != nil {
				fmt.Println("Error ", err)

			} else {
				fmt.Println("Success ", result.StartingInstances)
			}
		}

	}

}
