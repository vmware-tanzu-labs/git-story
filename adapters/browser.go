package adapters

import (
	"os/exec"

	"github.com/vmware-tanzu-labs/git-story/usecases"
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
