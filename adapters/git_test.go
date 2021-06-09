package adapters_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/git-story-branch/git-story-branch/adapters"
)

var _ = Describe("Git", func() {
	path := os.Getenv("PATH")

	BeforeEach(func() {
		os.Setenv("PATH", fmt.Sprintf("./fixtures:%s", path))
	})
	AfterEach(func() {
		os.Setenv("PATH", path)
	})

	It("should know the current branch", func() {
		branchName := adapters.NewRepository().GetBranchName()
		Expect(branchName).To(Equal("current-branch-123456789"))
	})

	It("should get all branch names", func() {
		branchNames := adapters.NewRepository().GetAllBranchNames()
		Expect(branchNames).To(Equal([]string{"main", "some-branch-#123"}))
	})

	It("should delete a branch", func() {
		error := adapters.NewRepository().DeleteBranch("some-accepted-branch-123456789")
		Expect(error).To(BeNil())
	})

	It("should fail to delete when there is a nonexistent branch", func() {
		error := adapters.NewRepository().DeleteBranch("some-nonexistent-branch")
		Expect(error).NotTo(BeNil())
		Expect(error.Error()).To(Equal("error: branch 'some-nonexistent-branch' not found."))
	})
})
