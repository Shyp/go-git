package octokit

import (
	"testing"

	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/jingweno/go-sawyer/hypermedia"
	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/jingweno/go-sawyer/mediaheader"
	"github.com/bmizerany/assert"
)

func TestNewResult_Pageable(t *testing.T) {
	resp := &Response{MediaHeader: &mediaheader.MediaHeader{Relations: hypermedia.Relations{"next": hypermedia.Hyperlink("/path")}}}
	result := newResult(resp, nil)

	assert.Equal(t, "/path", string(*result.NextPage))
	assert.T(t, result.PrevPage == nil)
	assert.T(t, result.LastPage == nil)
	assert.T(t, result.FirstPage == nil)
}
