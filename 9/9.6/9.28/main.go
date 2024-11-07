package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "name",
			// Aliases: []string{"c"},//現在StringFlagにはAliasesがない
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}
