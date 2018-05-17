package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"net/http"
	"os"
	"github.com/nlopes/slack"
	"fmt"
	"log"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token := os.Getenv(`WEBHOOK`)
	api := slack.New(token)

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Text:    "some text",
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("DAQFR3Z27", "Some text", params)
	if err != nil {
		log.Printf("%s\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: `Success`,
		}, nil
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
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