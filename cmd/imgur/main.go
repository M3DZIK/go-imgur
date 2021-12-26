package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-imgur/cmd/imgur/config"
	"github.com/mkideal/cli"
)

func main() {
	err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(upload)).
		Run(os.Args[1:])

	if err != nil {
		fmt.Println(err)
	}
}

var root = &cli.Command{
	Desc: "Imgur CLI",
}

var help = cli.HelpCommand("Help Menu")

func CreateClient() imgur.Client {
	config := config.ParseConfig()

	return imgur.Client{
		HTTPClient: new(http.Client),
		Imgur: imgur.Imgur{
			ClientID: config.Client.ID,
		},
	}
}

type uploadT struct {
	cli.Helper
	Title string `cli:"title" usage:"image title"`
}

var upload = &cli.Command{
	Name: "upload",
	Desc: "Upload image to Imgur",
	Argv: func() interface{} {
		return new(uploadT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*uploadT)

		CreateClient()

		ctx.String("%s\n", argv.Title)

		return nil
	},
}
