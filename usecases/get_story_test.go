package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/vmware-tanzu-labs/git-story/usecases"
)

var _ = Describe("Git Tracker name translator", func() {
	Describe("when the current branch is a story ID", func() {
		It("should retrieve a Pivotal Tracker Story based on the current git branch name", func() {
			mockGitRepo := &MockGitRepository{currentBranchName: "Insert Branch Name Here-#1234567890"}
			mockTrackerReader := MockPivotalTrackerReader{}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(error).To(BeNil())
			Expect(story.Description).To(Equal("Description"))
		})
	})
	Describe("when the current branch is not a story ID", func() {
		It("should return an error", func() {
			mockGitRepo := &MockGitRepository{currentBranchName: "main"}
			mockTrackerReader := MockPivotalTrackerReader{}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(story).To(BeNil())
			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("please run in branch that contains a Pivotal Tracker Story ID"))
		})
	})
	Describe("when the Tracker API returns an error", func() {
		It("should return an error", func() {
			mockGitRepo := &MockGitRepository{currentBranchName: "Insert Branch Name Here-#1234567890"}
			mockTrackerReader := MockPivotalTrackerReader{isBroken: true}

			story, error := usecases.GetStory(mockGitRepo, mockTrackerReader)

			Expect(story).To(BeNil())
			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("unable to find that story"))
		})
	})
})
