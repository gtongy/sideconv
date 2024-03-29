package main

import (
	"os"

	"github.com/gtongy/sideconv/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sideconv"
	app.Usage = "selenium IDE .side file converter"
	app.Commands = []cli.Command{
		{
			Name:    "create-app",
			Aliases: []string{"ca"},
			Usage:   "create initial app. please input app name.",
			Action: func(c *cli.Context) error {
				command.CreateApp(c)
				return nil
			},
		},
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "convert exec. convert-settings dir input convert settings.",
			Action: func(c *cli.Context) error {
				command.Convert(c)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
