package main

import (
	"context"

	"net/http"
	"os"
	"os/signal"
	"time"

	"yamcha/pkg/linebot"
	pkgStorage "yamcha/pkg/storage"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
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

	// regiest APIs
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/callback", bot.CallbackHandle)

	// Start server
	go func() {
		port := os.Getenv("PORT")
		log.Infof("service run at port %s", port)
		if err := e.Start(port); err != nil {
			log.Warn("shutting down the server, error:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
