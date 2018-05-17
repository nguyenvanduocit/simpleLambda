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
	"io/ioutil"
	json2 "encoding/json"
	"time"
	"math/rand"
)

type Quote struct {
	Author string `json:"author"`
	AuthorAvatar string `json:"author_avatar"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}

type Bot struct {
	Api *slack.Client
	Quotes []*Quote
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

	quote, err := bot.getRandomQuote()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body: err.Error(),
		}, nil
	}

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Title: "Today Quote",
		AuthorName: quote.Author,
		Text:    quote.Content,
		ImageURL: quote.AuthorAvatar,
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
	}
	return bot.getActions()
}

func (bot *Bot)loadQUotes()error{
	bQuotes, err := ioutil.ReadFile("./quotes.json")
	if err != nil {
		return err
	}
	if err := json2.Unmarshal(bQuotes, &bot.Quotes); err != nil {
		return err
	}
	return nil
}

func (bot *Bot)getRandomQuote()(*Quote, error){
	if len(bot.Quotes) == 0 {
		if err:= bot.loadQUotes(); err != nil {
			return nil, err
		}
	}
	rand.Seed(time.Now().Unix())
	return bot.Quotes[rand.Intn(len(bot.Quotes))], nil
}

func main() {
	bot := &Bot{
		Api: slack.New(os.Getenv(`SLACK_TOKEN`)),
	}
	lambda.Start(bot.handler)
}