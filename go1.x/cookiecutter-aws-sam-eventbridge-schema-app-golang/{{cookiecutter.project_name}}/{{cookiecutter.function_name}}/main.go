package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"{{ cookiecutter.codegen_path }}"
)

func handler(context context.Context, event {{ cookiecutter.codegen_package_name }}.AWSEvent) ([]byte , error) {
	/*
		// aws-lambda-go library automatically unmarshals the Lambda event JSON to the argument type used by your handler function
		// you can use the commented out code below to marshal JSON ([]byte) into a strongly typed struct.
		event, err := {{ cookiecutter.codegen_package_name }}.UnmarshalEvent(inputStream)
		if err != nil {
			return nil, err
		}
	*/

	// Retrieve {{ cookiecutter.AWS_Schema_name }} detail from event
	detail := event.Detail
	fmt.Println("Reading {{ cookiecutter.AWS_Schema_name }}: ", detail)

	// Developers write your event-driven business logic code here!
	// Make updates to the event payload
	fmt.Println("{{ cookiecutter.function_name }} updating event of " + event.DetailType)
	event.SetDetailType("{{ cookiecutter.function_name }} updated event of " + event.DetailType)

	// Return event as stream for further processing
	return {{ cookiecutter.codegen_package_name }}.Marshal(event)

}

func main() {
	lambda.Start(handler)
}
