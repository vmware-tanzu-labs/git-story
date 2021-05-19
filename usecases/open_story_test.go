package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

var _ = Describe("Browser Opener use case", func() {
	Describe("when the current branch is a story ID", func() {
		It("should open the browser to the URL of the story", func() {
			mockGitRepo := &MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			browserSpy := &BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader{}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).To(BeNil())
			Expect(browserSpy.openedURL).To(Equal("https://story.com/1234567890"))
		})
	})
	Describe("when the current branch is not a story ID", func() {
		It("should return an error", func() {
			mockGitRepo := &MockGitRepository{branchName: "main"}
			browserSpy := &BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader{}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("please run in branch that contains a Pivotal Tracker Story ID"))
		})
	})
	Describe("when the Tracker API returns an error", func() {
		It("should return an error", func() {
			mockGitRepo := &MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			browserSpy := &BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader{isBroken: true}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("unable to find that story"))
		})
	})
})
