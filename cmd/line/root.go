package line

import (
	"context"
	"os"
	"os/signal"
	"time"

	"yamcha/internal/config"
	"yamcha/internal/httputil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// LineCmd execute the yamcha application
var LineCmd = &cobra.Command{
	Use:   "line",
	Short: "TODO: 我是 Yamcha 的短短介紹",
	Long:  `TODO: 我是 Yamcha 的長篇介紹`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		// JUST TEST
		// repo, err := InitApplication(ctx)
		// if err != nil {
		// 	logrus.Fatal(err)
		// }

		logrus.Println("yamcha init...")
		cfg := config.NewConfiguration()

		// create Echo web service
		e := httputil.NewEcho(cfg)
		err := initService(e, cfg)
		if err != nil {
			logrus.Panicln("failed to register Restful API...")
			return
		}

		// Start server
		go func() {
			port := cfg.Server.Port
			logrus.Infof("service run at port %s", port)
			if err := e.Start(port); err != nil {
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
		if err := e.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
		}
	},
}
