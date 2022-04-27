package awsexample

import (
	"fmt"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

/* References :

*  https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
*  https://docs.aws.amazon.com/sdk-for-go/api/aws/session/#NewSessionWithOptions
*  https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
*  https://stackoverflow.com/questions/56135354/missingregion-could-not-find-region-configuration-in-golang-and-aws-sns
*  https://stackoverflow.com/questions/58069458/missingregion-could-not-find-region-configuration-but-i-have-it-in-my-aw
*
 */

func ListS3() {

	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		// Profile: "profile",

		//Provide SDK Config options, such as Region.
		// Config: aws.Config{
		// 	Region: aws.String("us-east-1"),
		// },

		// Force enable Shared Config support
		// Take Both Credentials and REGION from Profile
		//export AWS_PROFILE="aws_profile1"   ; echo $AWS_PROFILE
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Printf("Error Occurred! While Creating Session")
	}

	//session := session.New()
	svc := s3.New(sess)
	input := s3.ListBucketsInput{}

	result, err := svc.ListBuckets(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result.Buckets)
}
