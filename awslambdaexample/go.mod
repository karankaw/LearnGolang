module awslambdaexample

go 1.17

require (
	github.com/aws/aws-lambda-go v1.31.1
	github.com/aws/aws-sdk-go v1.44.5
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace example.com/startec2 => ../startec2

replace example.com/stopec2 => ../stopec2
