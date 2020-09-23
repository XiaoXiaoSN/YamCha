package line

import (
	"context"
	"os"
	"os/signal"
	"time"

	"yamcha/internal/config"
	"yamcha/internal/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// LineCmd execute the yamcha application
var LineCmd = &cobra.Command{
	Use:   "line",
	Short: "TODO: 我是 Yamcha 的短短介紹",
	Long:  `TODO: 我是 Yamcha 的長篇介紹`,
	Run: func(cmd *cobra.Command, args []string) {
		// JUST TEST
		// _, _ = initApplication(context.Background())

		log.Println("yamcha init...")
		cfg := config.NewConfiguration()

		// create Echo web service
		e := http.NewEcho(cfg)
		err := initService(e, cfg)
		if err != nil {
			log.Panicln("failed to register Restful API...")
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
	},
}
