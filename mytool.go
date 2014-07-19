package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "MyTool"
	app.Usage = "print arguments"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		fmt.Println("Number of arguments is", len(c.Args()))
		fmt.Println("First argument is", c.Args()[0])
		fmt.Println("Second argument is", c.Args()[1])
	}
	app.Run(os.Args)
}
