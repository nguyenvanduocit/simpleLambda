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
	"github.com/gin-gonic/gin/json"
)

type Bot struct {
	Api *slack.Client
}

func (bot *Bot)getChannels()(events.APIGatewayProxyResponse, error) {
	channels, err := bot.Api.GetChannels(false)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, nil
	}

	bChannels, err := json.Marshal(channels)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(bChannels),
	}, nil
}

func (bot *Bot)getQuote()(events.APIGatewayProxyResponse, error){
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Text:    "some text",
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := bot.Api.PostMessage("DAQFR3Z27", "Some text", params)
	if err != nil {
		log.Printf("%s\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, nil
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: `Success`,
	}, nil
}

func (bot *Bot)getActions()(events.APIGatewayProxyResponse, error){
	bActions, _ := json.Marshal([]string{
		"sendMorningQuote",
		"sendSleepQuote",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(bActions),
	}, nil
}

func (bot *Bot)handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token := request.QueryStringParameters["token"]
	if token != os.Getenv(`REQUEST_SECRET_KEY`) {
		bMessage, _ := json.Marshal(map[string]string{
			"error": "You are not allowed to call this function",
		})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string(bMessage),
		}, nil
	}
	action := request.QueryStringParameters["action"]
	switch action {
	case "channels":
		return bot.getChannels()
	}
	return bot.getChannels()
}

func main() {
	bot := &Bot{
		Api: slack.New(os.Getenv(`SLACK_TOKEN`)),
	}
	lambda.Start(bot.handler)
}