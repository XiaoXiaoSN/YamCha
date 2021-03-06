package line

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"yamcha/cmd/helper"
	"yamcha/internal/config"
	"yamcha/pkg/delivery/api"
	"yamcha/pkg/delivery/linebot"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// LineCmd execute the yamcha application
var LineCmd = &cobra.Command{
	Use:   "line",
	Short: "TODO: 我是 Yamcha 的短短介紹",
	Long:  `TODO: 我是 Yamcha 的長篇介紹`,
	Run: func(cmd *cobra.Command, args []string) {
		defer helper.CmdRecover()
		ctx := cmd.Context()

		app, err := InitApplication(ctx)
		if err != nil {
			logrus.Panic(err)
		}
		api.SetRoutes(app.Echo, app.Controller)
		linebot.SetRoutes(app.Echo, app.LineBot)

		// Start server
		go func() {
			port := fmt.Sprintf(":%d", config.Config().Server.Port)

			logrus.Infof("service run at port %s", port)
			if err := app.Echo.Start(port); err != nil {
				logrus.Warn("shutting down the server, error:", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := app.Echo.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
		}
	},
}
