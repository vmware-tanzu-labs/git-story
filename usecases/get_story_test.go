package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

type MockGitRepository struct {
}

func (mockGitRepo MockGitRepository) GetBranchName() string {
	return "Insert Branch Name Here-#1234567890"
}

type MockPivotalTrackerReader struct {
}

func (mockTrackerReader MockPivotalTrackerReader) GetStory(storyID int) *usecases.Story {
	if storyID == 1234567890 {
		return &usecases.Story{
			ID:          1234567890,
			Description: "Description",
			State:       "delivered",
		}
	}
	return nil
}

var _ = Describe("Git Tracker name translator", func() {
	It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
		mockGitRepo := MockGitRepository{}
		mockTrackerReader := MockPivotalTrackerReader{}

		story := usecases.GetStory(mockGitRepo, mockTrackerReader)

		Expect(story.Description).To(Equal("Description"))
		Expect(story.State).To(Equal("delivered"))
	})
})
