package main

import (
	"log"

	"github.com/MedzikUser/go-imgur/cmd/imgur/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgur <cmd>",
	Short: "Imgur API Cli",
	Long:  "Cli for Imgur API",
}

func main() {
	rootCmd.AddCommand(commands.UploadCmd)
	rootCmd.AddCommand(commands.DeleteCmd)

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
