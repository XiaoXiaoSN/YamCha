package main

import (
	"context"

	"os"
	"os/signal"
	"time"

	pkgConfig "yamcha/internal/pkg/config"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("yamcha init...")
	cfg := pkgConfig.NewConfiguration("config.yml")

	// create Echo web service
	e := echo.New()
	err := initService(e, cfg)
	if err != nil {
		log.Panicln("failed to regiest Restful API...")
		return
	}

	// Start server
	go func() {
		port := cfg.Server.Port
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
