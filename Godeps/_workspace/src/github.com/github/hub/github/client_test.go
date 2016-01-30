package github

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/Shyp/go-git/Godeps/_workspace/src/github.com/octokit/go-octokit/octokit"
	"github.com/bmizerany/assert"
)

func TestClient_newOctokitClient(t *testing.T) {
	c := NewClient("github.com")
	cc := c.newOctokitClient(nil)
	assert.Equal(t, "https://api.github.com", cc.Endpoint.String())

	c = NewClient("github.corporate.com")
	cc = c.newOctokitClient(nil)
	assert.Equal(t, "https://github.corporate.com", cc.Endpoint.String())
}

func TestClient_FormatError(t *testing.T) {
	e := &octokit.ResponseError{
		Response: &http.Response{
			StatusCode: 401,
			Status:     "401 Not Found",
		},
	}

	err := FormatError("action", e)
	assert.Equal(t, "Error action: Not Found (HTTP 401)", fmt.Sprintf("%s", err))

	e = &octokit.ResponseError{
		Response: &http.Response{
			StatusCode: 422,
			Status:     "422 Unprocessable Entity",
		},
		Message: "error message",
	}
	err = FormatError("action", e)
	assert.Equal(t, "Error action: Unprocessable Entity (HTTP 422)\nerror message", fmt.Sprintf("%s", err))
}

func TestClient_warnExistenceOfRepo(t *testing.T) {
	project := &Project{
		Name:  "hub",
		Owner: "github",
		Host:  "github.com",
	}
	e := &octokit.ResponseError{
		Response: &http.Response{
			StatusCode: 404,
			Status:     "404 Not Found",
		},
		Message: "error message",
	}

	err := warnExistenceOfRepo(project, e)
	assert.Equal(t, "Are you sure that github.com/github/hub exists?", fmt.Sprintf("%s", err))
}

func TestAuthTokenNote(t *testing.T) {
	note, err := authTokenNote(1)
	assert.Equal(t, nil, err)

	reg := regexp.MustCompile("hub for (.+)@(.+)")
	assert.T(t, reg.MatchString(note))

	note, err = authTokenNote(2)
	assert.Equal(t, nil, err)

	reg = regexp.MustCompile("hub for (.+)@(.+) 2")
	assert.T(t, reg.MatchString(note))

}
