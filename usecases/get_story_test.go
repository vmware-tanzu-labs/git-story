package usecases_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

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
		}, nil
	}
	return nil, nil
}

var _ = Describe("Git Tracker name translator", func() {
	Describe("when the current branch is a story ID", func() {
		It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
			mockGitRepo := MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			mockTrackerReader := MockPivotalTrackerReader{}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(error).To(BeNil())
			Expect(story.Description).To(Equal("Description"))
			Expect(story.State).To(Equal("delivered"))
		})
	})
	Describe("when the current branch is not a story ID", func() {
		It("should return an error", func() {
			mockGitRepo := MockGitRepository{branchName: "main"}
			mockTrackerReader := MockPivotalTrackerReader{}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(story).To(BeNil())
			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("Please run in branch that contains a Pivotal Tracker Story ID"))
		})
	})
	Describe("when the Tracker API returns an error", func() {
		It("should return an error", func() {
			mockGitRepo := MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			mockTrackerReader := MockPivotalTrackerReader{isBroken: true}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(story).To(BeNil())
			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("unable to find that story"))
		})
	})
})
