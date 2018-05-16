package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"net/http"
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := os.Getenv(`WEBHOOK`)
	data := map[string]interface{}{
		"attachments": map[string]interface{}{
			"title": "A quote everyday",
			"author_name": "Drunk Friend",
			"text": getQuote(),
		},
	}
	jsonValue, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: resp.StatusCode,
		Body: `Success`,
	}, nil
}

func getQuote()string {
	return `abc`
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}