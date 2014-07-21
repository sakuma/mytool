package mytool

import (
	"fmt"
	"os"
	"path"

	"github.com/sakuma/mytool/lib/utils"
)

var vimDirPath string = path.Join(os.Getenv("HOME"), ".vim")

func InstallVimConf() {
	_, err := os.Stat(vimDirPath)
	if err == nil {
		fmt.Println("already exits!")
		return
	}
	err = utils.Execute("git", "clone", "git@github.com:sakuma/dot.vim.git", vimDirPath)
	if err != nil {
		fmt.Println("Error: 'git clone'")
		return
	}

	err = os.Chdir(vimDirPath)
	if err != nil {
		fmt.Println("Error: chdir to vim conf dir")
		return
	}

	err = utils.Execute("./vimconf_ctl", "install")
	if err != nil {
		fmt.Println("Error: 'vimconf_ctl install'")
		return
	}
}

func UpdateVimConf() {
	err := os.Chdir(vimDirPath)
	if err != nil {
		fmt.Println("Error: chdir to vim conf dir")
		return
	}

	err = utils.Execute("./vimconf_ctl", "update")
	if err != nil {
		fmt.Println("Error: 'vimconf_ctl update'")
		return
	}
}
