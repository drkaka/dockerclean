package main

import (
	"fmt"
	"os"

	"github.com/drkaka/dockerclean/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dockerclean"
	app.Usage = "Tools for cleaning up image tags in docker registry."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "link, l",
			Usage: "the registry link. (should also provide username:password if registry requires)",
		},
		cli.IntFlag{
			Name:  "timeout, t",
			Usage: "Timeout seconds for requests. 0 means no timeout.",
			Value: 10,
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.GlobalString("link") == "" {
			fmt.Println("missing registry link to request")
			os.Exit(1)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list all images",
			Action: func(c *cli.Context) error {
				link := c.GlobalString("link")
				timeout := c.GlobalInt("timeout")
				return cmd.ListCommand(link, timeout)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
