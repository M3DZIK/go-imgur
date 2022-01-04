package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/0xAX/notificator"
	"github.com/MedzikUser/go-imgur/cmd/imgur/utils"
	"github.com/spf13/cobra"
)

var CreditsCmd = &cobra.Command{
	Use:                   "credits <delete hash>",
	Short:                 "Rate Limits",
	DisableFlagsInUseLine: true,
	Example:               "imgur credits",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := utils.CreateClient()

		limit, _, err := client.RateLimits()
		if err != nil {
			utils.Notify.Push("Error!", err.Error(), "", notificator.UR_CRITICAL)
			log.Fatal("Error Get Rate Limits: " + err.Error())
		}

		fmt.Printf("Client Total Credits     - %d\n", limit.Data.ClientLimit)
		fmt.Printf("Client Remaining Credits - %d\n", limit.Data.ClientRemaining)
		fmt.Printf("User Total Credits       - %d\n", limit.Data.UserLimit)
		fmt.Printf("User Remaining Credits   - %d\n", limit.Data.UserRemaining)
		fmt.Printf("User Credits Reset       - %s\n", time.Unix(int64(limit.Data.UserReset), 0).Format("2006-01-02 15:04:05 -0700"))

		return nil
	},
}
