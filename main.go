package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/sakuma/mytool/lib/mytool"
)

func main() {
	app := cli.NewApp()
	app.Name = "MyTool"
	app.Usage = "My tool config setting manager"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		command := c.Args()[0]
		switch command {
		case "install":
			fmt.Println("All install")
		case "update":
			fmt.Println("All update")
		default:
			fmt.Println("Unknown command: ", command)
		}
	}
	app.Commands = []cli.Command{
		{
			Name:      "vim",
			ShortName: "v",
			Usage:     "vim conf",
			Action: func(c *cli.Context) {
				command := c.Args()[0]
				switch command {
				case "install":
					mytool.InstallVimConf()
				case "update":
					mytool.UpdateVimConf()
				default:
					fmt.Println("Unknown command: ", command)
				}
			},
		},
		{
			Name:      "zsh",
			ShortName: "z",
			Usage:     "zsh conf",
			Action: func(c *cli.Context) {
				command := c.Args()[0]
				switch command {
				case "install":
					mytool.InstallZshConf()
				case "update":
					mytool.UpdateZshConf()
				default:
					fmt.Println("Unknown command: ", command)
				}
			},
		},
	}
	app.Run(os.Args)
}
