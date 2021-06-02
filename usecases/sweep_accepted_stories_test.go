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
		usecases.SweepAcceptedStories(mockGitRepo, mockTrackerReader)

		// expect mockGitRepo.delete to have been called with ome-accepted-story-#123
		Expect(mockGitRepo.deletedBranches).To(Equal([]string{"some-accepted-story-#123"}))
	})

	It("should return an error if the branch is unable to be deleted", func() {
		mockGitRepo := &MockGitRepository{branchNames: []string{"main", "some-accepted-story-#123", "some-wip-story-#234"}, isError: true}
		mockTrackerReader := MockPivotalTrackerReader{}
		usecases.SweepAcceptedStories(mockGitRepo, mockTrackerReader)

		Expect(nil).To(Equal("next: support a nice summary when we delete something"))
	})
})
