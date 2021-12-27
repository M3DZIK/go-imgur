package webhook

import (
	"fmt"

	"github.com/MedzikUser/go-imgur/cmd/imgur/config"
	"github.com/gtuk/discordwebhook"
)

func Send(url string, deletehash string) error {
	config := config.ParseConfig()

	var username = config.Discord.Username
	var description = fmt.Sprintf("Delete Hash: ||%s||", deletehash)

	msg := discordwebhook.Message{
		Username: &username,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:       &url,
				Description: &description,
				Color:       &config.Discord.EmbedColor,
				Image:       &discordwebhook.Image{Url: &url},
			},
		},
	}

	return discordwebhook.SendMessage(config.Discord.URL, msg)
}
