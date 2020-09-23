package main

import (
	"os"
	"yamcha/cmd/line"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
