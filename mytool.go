package main

import (
	"fmt"
	"os"
	"os/exec"
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
	_, err := os.Stat(vimDirPath)
	if err == nil {
		fmt.Println("already exits!")
		return
	}
	cmd := exec.Command("git", "clone", "git@github.com:sakuma/dot.vim.git", vimDirPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		fmt.Println("Error: 'git clone'")
		return
	}

	err = os.Chdir(vimDirPath)
	if err != nil {
		fmt.Println("Error: chdir to vim conf dir")
		return
	}

	cmd = exec.Command("./vimconf_ctl", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		fmt.Println("Error: 'vimconf_ctl install'")
		return
	}
}
