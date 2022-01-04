package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/0xAX/notificator"
	"github.com/MedzikUser/go-imgur/cmd/imgur/utils"
	u "github.com/MedzikUser/go-utils/utils"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:                   "info <image id>",
	Short:                 "Image Info",
	Long:                  "Get Image Info",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Example:               "imgur info n744BL9",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := utils.CreateClient()

		data, _, err := client.GetImageInfo(args[0])
		if err != nil {
			utils.Notify.Push("Error!", err.Error(), "", notificator.UR_CRITICAL)
			log.Fatal("Error get image info: " + err.Error())
		}

		fmt.Printf("ID          - %s\n", data.Data.ID)
		fmt.Printf("URL         - %s\n", "https://cdn.magicuser.cf/"+data.Data.IDExt)
		fmt.Printf("Views       - %d\n", data.Data.Views)
		fmt.Printf("Size        - %s\n", u.Bytes(uint64(data.Data.Size)))
		fmt.Printf("Height      - %d\n", data.Data.Height)
		fmt.Printf("Width       - %d\n", data.Data.Width)
		fmt.Printf("Type        - %s\n", data.Data.MimeType)
		fmt.Printf("Upload Date - %s\n", time.Unix(int64(data.Data.Datetime), 0).Format("2006-01-02 15:04:05 -0700"))
		fmt.Printf("Bandwidth   - %s\n", u.Bytes(uint64(data.Data.Bandwidth)))

		return nil
	},
}
