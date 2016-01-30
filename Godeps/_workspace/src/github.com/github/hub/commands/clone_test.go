package commands

import (
	"os"
	"testing"

	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/github/hub/github"
	"github.com/bmizerany/assert"
)

func TestTransformCloneArgs(t *testing.T) {
	os.Setenv("HUB_PROTOCOL", "git")
	github.CreateTestConfigs("jingweno", "123")

	args := NewArgs([]string{"clone", "foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git://github.com/foo/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:foo/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "jingweno/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git://github.com/jingweno/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "acl-services/devise-acl"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:acl-services/devise-acl.git", args.FirstParam())

	args = NewArgs([]string{"clone", "jekyll_and_hyde"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git://github.com/jingweno/jekyll_and_hyde.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "jekyll_and_hyde"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:jingweno/jekyll_and_hyde.git", args.FirstParam())

	args = NewArgs([]string{"clone", "git://github.com/jingweno/gh", "gh"})
	transformCloneArgs(args)

	assert.Equal(t, 2, args.ParamsSize())
	assert.Equal(t, "git://github.com/jingweno/gh", args.FirstParam())
	assert.Equal(t, "gh", args.GetParam(1))
}
