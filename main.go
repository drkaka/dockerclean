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
		{
			Name:  "tags",
			Usage: "list all tags for an image",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image, i",
					Usage: "the image to search",
				},
			},
			Action: func(c *cli.Context) error {
				link := c.GlobalString("link")
				timeout := c.GlobalInt("timeout")
				image := c.String("image")
				if len(image) == 0 {
					fmt.Println("please specify an image")
					os.Exit(1)
				}
				return cmd.TagsCommand(link, image, timeout)
			},
		},
		{
			Name:  "delete",
			Usage: "delete the given tag of the image",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image, i",
					Usage: "the image to search",
				},
				cli.StringFlag{
					Name:  "tag, t",
					Usage: "the tag of the image",
				},
			},
			Action: func(c *cli.Context) error {
				link := c.GlobalString("link")
				timeout := c.GlobalInt("timeout")
				image := c.String("image")
				if len(image) == 0 {
					fmt.Println("please specify an image")
					os.Exit(1)
				}
				tag := c.String("tag")
				if len(tag) == 0 {
					fmt.Println("please specify a tag")
					os.Exit(1)
				}
				return cmd.DeleteTagCommand(link, image, tag, timeout)
			},
		},
		{
			Name:  "keep",
			Usage: "keep the latest N tags",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image, i",
					Usage: "the image to search",
				},
				cli.IntFlag{
					Name:  "number, n",
					Usage: "the number of tags to keep",
					Value: 5,
				},
			},
			Action: func(c *cli.Context) error {
				link := c.GlobalString("link")
				timeout := c.GlobalInt("timeout")
				image := c.String("image")
				if len(image) == 0 {
					fmt.Println("please specify an image")
					os.Exit(1)
				}
				num := c.Int("number")
				if num <= 0 {
					fmt.Println("please specify a number larger than 0")
					os.Exit(1)
				}
				return cmd.KeepCommand(link, image, num, timeout)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
