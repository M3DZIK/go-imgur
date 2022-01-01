package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/0xAX/notificator"
	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-imgur/cmd/imgur/config"
	"github.com/MedzikUser/go-imgur/cmd/imgur/utils"
	"github.com/MedzikUser/go-imgur/cmd/imgur/webhook"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:                   "upload <path to file or url>",
	Short:                 "Upload image to Imgur",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Example:               "imgur upload Pictures/Screenshot.png",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := utils.CreateClient()

		f, err := os.Stat(args[0])

		if err == nil {
			if f.IsDir() {
				log.Fatalf("%s is dir!", args[0])
			}

			data, _, err := client.UploadImageFromFile(args[0], "")
			if err != nil {
				log.Fatal("Error upload image to Imgur: " + err.Error())
			}

			printLink(data)
		} else {
			data, _, err := client.UploadImageFromURL(args[0], "")
			if err != nil {
				log.Fatal("Error upload image to Imgur: " + err.Error())
			}

			printLink(data)
		}

		return nil
	},
}

func printLink(data *imgur.ImageInfoData) {
	url := "https://cdn.magicuser.cf/" + data.Data.IDExt

	fmt.Printf("URL          %s\n", url)
	fmt.Printf("Delete Hash  %s\n", data.Data.Deletehash)

	err := clipboard.WriteAll(url)
	if err != nil {
		fmt.Println("Error copy link to clipboard: " + err.Error())
	}

	notify := notificator.New(notificator.Options{
		AppName: "Imgur",
	})

	notify.Push("Uploaded!", url, "", notificator.UR_NORMAL)

	config := config.ParseConfig()
	if config.Discord.Enable {
		webhook.Send(url, data.Data.Deletehash)
	}
}
