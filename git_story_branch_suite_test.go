package storybranch_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGitStoryBranch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitStoryBranch Suite")
}
