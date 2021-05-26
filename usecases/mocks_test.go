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
	branchName      string
	branchNames     []string
	deletedBranches []string
}

// TODO rename this to GetCurrentBranchName()
func (mockGitRepo MockGitRepository) GetBranchName() string {
	return mockGitRepo.branchName
}

func (mockGitRepo MockGitRepository) GetAllBranchNames() []string {
	return mockGitRepo.branchNames
}

func (mockGitRepo *MockGitRepository) DeleteBranch(branchName string) *exec.Cmd {
	mockGitRepo.deletedBranches = append(mockGitRepo.deletedBranches, branchName)
	return nil
}

type MockPivotalTrackerReader struct {
	isBroken bool
}

func (mockTrackerReader MockPivotalTrackerReader) GetStory(storyID int) (*usecases.Story, error) {
	if mockTrackerReader.isBroken {
		return nil, errors.New("unable to find that story")
	}
	if storyID == 123 {
		return &usecases.Story{
			ID:          123,
			Description: "Description",
			State:       "accepted",
			URL:         "https://story.com/123",
		}, nil
	}
	if storyID == 234 {
		return &usecases.Story{
			ID:          234,
			Description: "Description",
			State:       "started",
			URL:         "https://story.com/234",
		}, nil
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
