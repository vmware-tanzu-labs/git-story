package usecases

import (
	"os/exec"
)

func OpenStory(repo Repository, tracker Tracker, browser Browser) error {
	story, error := GetStory(repo, tracker)
	if error != nil {
		return error
	}
	_, error = browser.OpenURL(story.URL)
	return error
}

type Browser interface {
	OpenURL(URL string) (*exec.Cmd, error)
}
