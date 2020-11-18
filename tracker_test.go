package storybranch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	storybranch "github.com/carpeliam/git-story-branch"
)

var _ = Describe("Tracker", func() {

	It("should be able to look up the description of a story", func() {
		// arrange
		// some kind of interface
		// maybe we need to mock out tracker to have a story id of 123456789

		// act
		description := storybranch.GetStoryDescription(123456789)

		// assert
		Expect(description).To(Equal("I dunno, uh, cool story... bro.. or something."))
	})
})
