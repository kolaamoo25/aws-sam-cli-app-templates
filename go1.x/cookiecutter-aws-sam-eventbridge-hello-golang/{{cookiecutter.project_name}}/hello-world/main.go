package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"{{ cookiecutter.project_name }}/hello-world/model/aws/ec2"
	"{{ cookiecutter.project_name }}/hello-world/model/aws/ec2/marshaller"
)

func handler(context context.Context, awsEvent ec2.AWSEvent) ([]byte , error) {
	/*
		// aws-lambda-go library automatically unmarshals the Lambda event JSON to the argument type used by your handler function
		// you can use the commented out code below to convert JSON ([]byte) into a strongly typed struct
		event, err := marshaller.UnmarshalEvent(inputStream)
		if err != nil {
			return nil, err
		}
	*/

	// Retrieve the ec2 notification from the event
	ec2InstanceStateChangeNotification := awsEvent.Detail

	// Developers write your event-driven business logic code here!
	fmt.Println("Instance " + ec2InstanceStateChangeNotification.InstanceId + " transitioned to " + ec2InstanceStateChangeNotification.State)

	// Make updates to the event payload
	awsEvent.SetDetailType("HelloWorldFunction updated event of " + awsEvent.DetailType)

	// Return event as stream for further processing
	return marshaller.Marshal(awsEvent)

}

func main() {
	lambda.Start(handler)
}
