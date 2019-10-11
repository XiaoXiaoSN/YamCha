package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"yamcha/pkg/linebot"
	pkgStorage "yamcha/pkg/storage"
)

var bot linebot.LineBot

func main() {
	var err error

	storage := pkgStorage.NewMemoryStorage()

	channelSecret := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	channelToken := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")
	if bot, err = linebot.NewYambotLineBot(channelSecret, channelToken, storage); err != nil {
		log.Println("Bot:", bot, " err:", err)
		return
	}

	// regiest BOT APIs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/callback", bot.CallbackHandle)

	// run http service
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Http Service run at port %s\n", addr)
	if err = http.ListenAndServe(addr, nil); err != nil {
		log.Println("End Http Service...")
	}
}
