package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-imgur/cmd/imgur/utils"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:   "upload <path to file or url>",
	Short: "Upload image to Imgur",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := utils.CreateClient()

		f, err := os.Stat(args[0])

		if err == nil {
			if f.IsDir() {
				log.Fatalf("%s is dir", args[0])
			}

			data, _, err := client.UploadImageFromFile(args[0], "")
			if err != nil {
				log.Fatal(err)
			}

			printLink(data)
		} else {
			data, _, err := client.UploadImageFromURL(args[0], "")
			if err != nil {
				log.Fatal(err)
			}

			printLink(data)
		}

		return nil
	},
}

func printLink(data *imgur.ImageInfoData) {
	fmt.Printf("URL          %s\n", data.Data.Link)
	fmt.Printf("Delete Hash  %s\n", data.Data.Deletehash)
}
