package main

import (
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "config",
			// Aliases: []string{"c"},
			Usage: "Load configration from `FILE`",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "output as JSON",
					Value: false,
				},
			},
			Action: cmdList,
		},
	}

}
