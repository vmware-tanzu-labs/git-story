package usecases_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

type BrowserSpy struct {
	openedURL string
}

func (browserSpy BrowserSpy) OpenURL(URL string) error {
	browserSpy.openedURL = URL
	return nil
}

type MockGitRepository2 struct {
	branchName string
}

func (mockGitRepo MockGitRepository2) GetBranchName() string {
	return mockGitRepo.branchName
}

type MockPivotalTrackerReader2 struct {
	isBroken bool
}

func (mockTrackerReader MockPivotalTrackerReader2) GetStory(storyID int) (*usecases.Story, error) {
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

var _ = Describe("Browser Opener use case", func() {
	Describe("when the current branch is a story ID", func() {
		It("should open the browser to the URL of the story", func() {
			mockGitRepo := MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			browserSpy := BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader2{}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).To(BeNil())
			Expect(browserSpy.openedURL).To(Equal("https://story.com/1234567890"))
		})
	})
	Describe("when the current branch is not a story ID", func() {
		It("should return an error", func() {
			mockGitRepo := MockGitRepository{branchName: "main"}
			browserSpy := BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader{}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("Please run in branch that contains a Pivotal Tracker Story ID"))
		})
	})
	Describe("when the Tracker API returns an error", func() {
		It("should return an error", func() {
			mockGitRepo := MockGitRepository{branchName: "Insert Branch Name Here-#1234567890"}
			browserSpy := BrowserSpy{}
			mockTrackerReader := MockPivotalTrackerReader{isBroken: true}

			error := usecases.OpenStory(mockGitRepo, mockTrackerReader, browserSpy)

			Expect(error).NotTo(BeNil())
			Expect(error.Error()).To(ContainSubstring("unable to find that story"))
		})
	})
})
