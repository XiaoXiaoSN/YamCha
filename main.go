package main

import (
	"os"
	"yamcha/cmd/line"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "gorm.io/driver/mysql"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "choose instance to run: [ line ]",
}

func main() {
	rootCmd.AddCommand(line.LineCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Errorf("rootCmd.Execute failed. err: %+v", err)
		os.Exit(-1)
	}
}
