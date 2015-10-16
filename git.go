package git

import (
	"os/exec"
	"strings"
)

func CurrentBranch() (string, error) {
	result, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}
