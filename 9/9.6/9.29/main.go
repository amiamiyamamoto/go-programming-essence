package main

import (
	"encoding/json"
	"fmt"
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
		return listStudentsAsJSON()
	}
	return listStudent()
}

func listStudentsAsJSON() error {
	score := []struct {
		Name  string
		Score int
	}{
		{
			Name:  "Alice",
			Score: 100,
		},
		{
			Name:  "Bob",
			Score: 90,
		},
	}
	jsondata, err := json.Marshal(score)
	if err != nil {
		return err
	}
	fmt.Println(string(jsondata))
	return nil
}

func listStudent() error {
	fmt.Println("Alice: 100")
	fmt.Println("Bob: 90")
	return nil
}
