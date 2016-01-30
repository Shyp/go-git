package commands

import (
	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/github/hub/github"
	"github.com/bmizerany/assert"
	"testing"
)

func TestTransformSubmoduleArgs(t *testing.T) {
	github.CreateTestConfigs("jingweno", "123")

	args := NewArgs([]string{"submodule", "add", "jingweno/gh", "vendor/gh"})
	transformSubmoduleArgs(args)

	cmds := args.Commands()
	assert.Equal(t, 1, len(cmds))
	assert.Equal(t, "git submodule add git://github.com/jingweno/gh.git vendor/gh", cmds[0].String())

	args = NewArgs([]string{"submodule", "add", "-p", "jingweno/gh",
		"vendor/gh"})
	transformSubmoduleArgs(args)

	cmds = args.Commands()
	assert.Equal(t, 1, len(cmds))
	assert.Equal(t, "git submodule add git@github.com:jingweno/gh.git vendor/gh", cmds[0].String())

	args = NewArgs([]string{"submodule", "add", "-b", "gh", "--name", "gh", "jingweno/gh", "vendor/gh"})
	transformSubmoduleArgs(args)

	cmds = args.Commands()
	assert.Equal(t, 1, len(cmds))
	assert.Equal(t, "git submodule add -b gh --name gh git://github.com/jingweno/gh.git vendor/gh", cmds[0].String())
}
