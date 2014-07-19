package main

import (
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "MyTool"
	app.Usage = "print arguments"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		command := c.Args()[0]
		fmt.Println("First argument is", command)
		switch command {
		case "install":
			setupVim()
		case "update":
			fmt.Println("update...", command)
		default:
			fmt.Println("Unknown command: ", command)
		}
	}
	app.Run(os.Args)
}

func setupVim() {
	vimDirPath := path.Join(os.Getenv("HOME"), ".vim")
	if _, err := os.Stat(vimDirPath); err == nil {
		fmt.Println("exits!")
	} else {
		fmt.Println("not exits!")
	}
}
