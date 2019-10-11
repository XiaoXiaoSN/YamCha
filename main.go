package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	// serverURL := os.Getenv("LINECORP_PLATFORM_CHANNEL_SERVERURL")
	channelSecret := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	channelToken := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")
	if bot, err = linebot.New(channelSecret, channelToken); err != nil {
		log.Println("Bot:", bot, " err:", err)
		return
	}

	log.Println("Channel Secret:", channelSecret)
	log.Println("Channel Token:", channelToken)

	// BOT APIs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/callback", callbackHandler)

	// provide by Heroku
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Http Service run at port %s\n", addr)
	if err = http.ListenAndServe(addr, nil); err != nil {
		log.Println("End Http Service...")
	}
}
