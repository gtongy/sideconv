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
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "convert exec",
			Action: func(c *cli.Context) error {
				command.Convert(c)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
