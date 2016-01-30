package github

import (
	"os"

	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/github/hub/git"
	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/mattn/go-isatty"
)

func IsHttpsProtocol() bool {
	httpProcotol, _ := git.Config("hub.protocol")
	if httpProcotol == "https" {
		return true
	}

	httpClone, _ := git.Config("--bool hub.http-clone")
	if httpClone == "true" {
		return true
	}

	return false
}

func IsTerminal(f *os.File) bool {
	return isatty.IsTerminal(f.Fd())
}
