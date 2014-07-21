package mytool

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func InstallVimConf() {
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

func UpdateVimConf() {
	vimDirPath := path.Join(os.Getenv("HOME"), ".vim")
	err := os.Chdir(vimDirPath)
	if err != nil {
		fmt.Println("Error: chdir to vim conf dir")
		return
	}

	cmd := exec.Command("./vimconf_ctl", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		fmt.Println("Error: 'vimconf_ctl update'")
		return
	}
}
