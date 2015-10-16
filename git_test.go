package git

import (
	"fmt"
	"testing"
)

func TestCurrentBranch(t *testing.T) {
	result, err := CurrentBranch()
	fmt.Println(err)
	fmt.Println(result)
	fmt.Println(len(result))
}
