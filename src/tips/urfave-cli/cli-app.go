package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "List students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "Output as JSON",
					Value: false,
				},
			},
			Action: func(c *cli.Context) error {
				if c.Bool("json") {
					// do something
					return nil
				}
				return nil
			},
		},
	}

	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}
