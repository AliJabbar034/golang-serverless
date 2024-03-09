package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event *MyEvent) (*Response, error) {

	return &Response{
		Message: fmt.Sprintf("my name is %s", event.Name),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
