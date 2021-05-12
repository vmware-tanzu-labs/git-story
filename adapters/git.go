package adapters

import (
	"os/exec"
	"strings"
)

// GitRepository is an abstraction for git repositories
type GitRepository struct{}

// GetBranchName gets the current branch name of the repository.
func (repo GitRepository) GetBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(output))
}

func (repo GitRepository) DeleteBranch(branchName string) {
}
func (repo GitRepository) GetAllBranchNames() []string {
	return []string{}
}

// NewRepository creates  a new repository
func NewRepository() *GitRepository {
	return &GitRepository{}
}
