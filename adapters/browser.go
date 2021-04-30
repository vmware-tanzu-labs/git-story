package adapters

import (
	"fmt"
	"os/exec"

	"github.com/git-story-branch/git-story-branch/usecases"
)

type Browser struct{}

func (browser Browser) OpenURL(url string) (*exec.Cmd, error) {
	fmt.Println("url")
	fmt.Println(url)
	cmd := exec.Command("open", url)
	error := cmd.Start()
	return cmd, error
	// return .Start()
}

func NewBrowser() usecases.Browser {
	return &Browser{}
}
