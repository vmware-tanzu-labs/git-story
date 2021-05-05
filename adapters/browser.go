package adapters

import (
	"os/exec"

	"github.com/git-story-branch/git-story-branch/usecases"
)

type Browser struct{}

func (browser Browser) OpenURL(url string) (*exec.Cmd, error) {
	cmd := exec.Command("open", url)
	error := cmd.Start()
	return cmd, error
}

func NewBrowser() usecases.Browser {
	return &Browser{}
}
