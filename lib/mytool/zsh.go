package mytool

import (
	"fmt"
	"os"
	"path"

	"../utils"
)

// var zshDirPath string = path.Join(os.Getenv("HOME"), ".zsh.d")
var zshDirPath string = path.Join("/tmp", "zsh.d")

func InstallZshConf() {
	_, err := os.Stat(zshDirPath)
	if err == nil {
		fmt.Println("already exits!")
		return
	}
	err = utils.Execute("git", "clone", "git@github.com:sakuma/zsh.d.git", zshDirPath)
	if err != nil {
		fmt.Println("Error: 'git clone'")
		return
	}

	err = os.Chdir(zshDirPath)
	if err != nil {
		fmt.Println("Error: chdir to zsh conf dir")
		return
	}

	err = utils.Execute("ln", "-fs", (zshDirPath + "/dot.zprofile"), (os.Getenv("HOME") + "/.zprofile"))
	if err != nil {
		return
	}

	err = utils.Execute("cp", (zshDirPath + "/conf/switch_view.zsh.example"), (zshDirPath + "/conf/switch_view.zsh"))
	if err != nil {
		return
	}
}

func UpdateZshConf() {
	err := os.Chdir(zshDirPath)
	if err != nil {
		fmt.Println("Error: chdir to zsh conf dir")
		return
	}

	err = utils.Execute("git", "pull", "origin", "master")
	if err != nil {
		fmt.Println("Error: 'git pull origin master' on zsh conf")
		return
	}
}
