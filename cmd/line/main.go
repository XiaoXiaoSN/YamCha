package main

import (
	"context"

	"net/http"
	"os"
	"os/signal"
	"time"

	pkgDB "yamcha/internal/pkg/database"
	"yamcha/pkg/linebot"

	pkgOrder "yamcha/pkg/api/order"
	orderRepo "yamcha/pkg/api/order/repository"
	orderSvc "yamcha/pkg/api/order/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

var bot linebot.LineBot

func main() {
	var err error

	channelSecret := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	channelToken := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")

	// need modify
	var (
		_orderRepo pkgOrder.Repository
		_orderSvc  pkgOrder.Service
	)
	db, err := pkgDB.NewDatabases(pkgDB.Config{
		Username: "xiao",
		Password: "gUKmFVmSdOgTTinmQa9fmYr5AT0EAci5",
		Address:  "yamcha.10oz.tw:23306",
		DBName:   "yamcha_db",
		Env:      "dev",
	})
	if err != nil {
		log.Println("err:", err)
	}
	_orderRepo = orderRepo.NewOrderRepository(db)
	_orderSvc = orderSvc.NewOrderService(_orderRepo)

	if bot, err = linebot.NewYambotLineBot(channelSecret, channelToken, _orderSvc); err != nil {
		log.Println("Bot:", bot, " err:", err)
		return
	}
	// need modify

	// regiest APIs
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middlewareCfg))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/callback", bot.CallbackHandle)

	err = initRestfulAPI(e)
	if err != nil {
		log.Panicln("failed to regiest Restful API...")
		return
	}

	// Start server
	go func() {
		port := ":" + os.Getenv("PORT")
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
