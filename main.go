package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var serverURL string

func main() {
	var err error
	serverURL = os.Getenv("LINECORP_PLATFORM_CHANNEL_SERVERURL")
	channelSecret := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	channelToken := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")
	if bot, err = linebot.New(channelSecret, channelToken); err != nil {
		log.Println("Bot:", bot, " err:", err)
		return
	}

	log.Println("Channel Secret:", channelSecret)
	log.Println("Channel Token:", channelToken)

	//BOT APIs
	http.HandleFunc("/callback", callbackHandler)

	//provide by Heroku
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}
