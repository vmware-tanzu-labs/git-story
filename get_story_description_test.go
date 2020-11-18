package storybranch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	storybranch "github.com/carpeliam/git-story-branch"
)

type MockGitRepository struct {
}

func (mockGitRepo MockGitRepository) GetBranchName() string {
	return "Insert Branch Name Here-#1234567890"
}

type MockPivotalTrackerReader struct {
}

func (mockTrackerReader MockPivotalTrackerReader) GetStoryDescription(storyID int) string {
	if storyID == 1234567890 {
		return "Description"
	}
	return "Wrong Description"
}

var _ = Describe("Git Tracker name translator", func() {
	It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
		mockGitRepo := MockGitRepository{}
		mockTrackerReader := MockPivotalTrackerReader{}

		storyDescription := storybranch.GetStoryDescription(mockGitRepo, mockTrackerReader)

		Expect(storyDescription).To(Equal("Description"))
	})
})
