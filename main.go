package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"net/http"
	"os"
	"github.com/nlopes/slack"
	"fmt"
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

	bChannels, _ := json.Marshal(channels)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(bChannels),
	}, nil
}

func (bot *Bot)sendQuote(channelId string)(events.APIGatewayProxyResponse, error){
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Title: "Today Quote",
		Text:    "some text",
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := bot.Api.PostMessage(channelId, "", params)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, nil
	}

	bMessage, _ := json.Marshal(map[string]string{
		"message": fmt.Sprintf("Message successfully sent to channel %s at %s", channelID, timestamp),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: string(bMessage),
	}, nil
}

func (bot *Bot)getActions()(events.APIGatewayProxyResponse, error){
	bActions, _ := json.Marshal([]string{
		"channels",
		"morningQuote",
		"sleepingQuote",
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
	secret := request.QueryStringParameters["secret"]
	if secret != os.Getenv(`REQUEST_SECRET`) {
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
	case "morningQuote":
		channelId := request.QueryStringParameters["channelId"]
		return bot.sendQuote(channelId)
	case "sleepingQuote":
		channelId := request.QueryStringParameters["channelId"]
		return bot.sendQuote(channelId)
	}
	return bot.getActions()
}

func main() {
	bot := &Bot{
		Api: slack.New(os.Getenv(`SLACK_TOKEN`)),
	}
	lambda.Start(bot.handler)
}