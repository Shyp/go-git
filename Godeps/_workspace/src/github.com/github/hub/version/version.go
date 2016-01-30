package version

import (
	"fmt"

	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/github/hub/git"
	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/github/hub/utils"
)

var Version = "2.2.2"

func FullVersion() string {
	gitVersion, err := git.Version()
	utils.Check(err)
	return fmt.Sprintf("%s\nhub version %s", gitVersion, Version)
}
