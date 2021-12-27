package commands

import (
	"fmt"
	"log"

	"github.com/MedzikUser/go-imgur/cmd/imgur/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <delete hash>",
	Short: "Delete image from Imgur by Delete Hash",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := utils.CreateClient()

		_, _, err := client.DeleteImageUnAuthed(args[0])
		if err != nil {
			log.Fatal("Error delete image: " + err.Error())
		}

		fmt.Println("Deleted image!")

		return nil
	},
}
