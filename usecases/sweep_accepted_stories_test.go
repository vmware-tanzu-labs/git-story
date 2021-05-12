package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	usecases "github.com/git-story-branch/git-story-branch/usecases"
)

var _ = Describe("Story Sweeper use case", func() {
	It("should delete all branches that correspond to accepted stories", func() {
		mockGitRepo := &MockGitRepository{branchNames: []string{"main", "some-accepted-story-#123", "some-wip-story-#234"}}
		mockTrackerReader := MockPivotalTrackerReader{}
		result := usecases.SweepAcceptedStories(mockGitRepo, mockTrackerReader)

		var resultContainsStory bool
		_, resultContainsStory = result["some-accepted-story-#123"]

		Expect(resultContainsStory).To(BeTrue())
		Expect(mockGitRepo.deletedBranches).To(Equal([]string{"some-accepted-story-#123"}))
	})

	It("should return an error if the branch is unable to be deleted", func() {
		mockGitRepo := &MockGitRepository{
			branchNames:     []string{"main", "some-accepted-story-#123", "some-accepted-story-#789", "some-wip-story-#234"},
			erroredBranches: []string{"some-accepted-story-#123"},
		}
		mockTrackerReader := MockPivotalTrackerReader{}
		results := usecases.SweepAcceptedStories(mockGitRepo, mockTrackerReader)

		Expect(results["some-accepted-story-#123"]).NotTo(BeNil())
		Expect(results["some-accepted-story-#789"]).To(BeNil())
	})
})
