package main

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
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
	app.Name = "Score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}

func cmdList(c *cli.Context) error {
	if c.Bool("json") {
		return nil
	}
	return nil
}
