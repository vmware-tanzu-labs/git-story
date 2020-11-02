package storybranch_test

import (
	storybranch "github.com/carpeliam/git-story-branch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GitStoryBranch", func() {
	It("should know about stories and their names", func() {
		storyName := storybranch.GetSpikeStoryName()
		Expect(storyName).To(Equal("[SPIKE] Validate we can talk to git and Tracker"))
	})

	It("should know about branches", func() {
		branchName := storybranch.GetBranchName()
		Expect(branchName).To(Equal("SPIKE-use-git-and-tracker-175526301"))
	})
})
