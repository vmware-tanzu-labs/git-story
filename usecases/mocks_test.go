package usecases_test

import (
	"errors"
	"os/exec"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

type BrowserSpy struct {
	openedURL string
}

func (browserSpy *BrowserSpy) OpenURL(URL string) (*exec.Cmd, error) {
	browserSpy.openedURL = URL
	return nil, nil
}

type MockGitRepository struct {
	branchName string
}

func (mockGitRepo MockGitRepository) GetBranchName() string {
	return mockGitRepo.branchName
}

type MockPivotalTrackerReader struct {
	isBroken bool
}

func (mockTrackerReader MockPivotalTrackerReader) GetStory(storyID int) (*usecases.Story, error) {
	if mockTrackerReader.isBroken {
		return nil, errors.New("unable to find that story")
	}
	if storyID == 1234567890 {
		return &usecases.Story{
			ID:          1234567890,
			Description: "Description",
			State:       "delivered",
			URL:         "https://story.com/1234567890",
		}, nil
	}
	return nil, nil
}
