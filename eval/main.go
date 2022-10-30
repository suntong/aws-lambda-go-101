package main

import (
	"log"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	log.Printf("GET: %v", req)
	cmd := req.QueryStringParameters["cmd"]
	log.Printf("Cmd: '%s'", cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Printf("Failed to execute command: %s", cmd)
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(out),
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
