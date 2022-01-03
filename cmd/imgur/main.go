package main

import (
	"log"

	"github.com/MedzikUser/go-imgur/cmd/imgur/commands"
	"github.com/spf13/cobra"
)

var version = "unknow"

var rootCmd = &cobra.Command{
	Use:     "imgur <cmd>",
	Short:   "Imgur API CLI",
	Long:    "CLI for Imgur API",
	Version: version,
}

func main() {
	rootCmd.AddCommand(commands.UploadCmd)
	rootCmd.AddCommand(commands.InfoCmd)
	rootCmd.AddCommand(commands.DeleteCmd)

	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
