package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

//LinkCustomer : A chatbot DB to store account link information.
type LinkCustomer struct {
	//Data from CustData from provider.
	Name   string
	Age    int
	Desc   string
	Nounce string
	//For chatbot linked data.
	LinkUserID string
}

var linkedCustomers []LinkCustomer

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch {
				case strings.EqualFold(message.Text, "快速回復按鈕"):
					_ = quickReplyButton(bot, event)
					return
				case strings.EqualFold(message.Text, "t"):
					date := "20191010"
					res, err := bot.GetNumberMulticastMessages(date).Do()
					fmt.Println(res, err)
					return
				}
			}

		} else if event.Type == linebot.EventTypeAccountLink {
			log.Println("event.Type == linebot.EventTypeAccountLink")
		}
	}
}

func quickReplyButton(bot *linebot.Client, event *linebot.Event) error {
	_, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("Welcome").
			WithQuickReplies(linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton(
					"",
					linebot.NewMessageAction("account link", "link")),
				linebot.NewQuickReplyButton(
					"",
					linebot.NewMessageAction("list user", "list")),
			)),
	).Do()
	if err != nil {
		log.Println("err:", err)
		return err
	}

	return nil
}
