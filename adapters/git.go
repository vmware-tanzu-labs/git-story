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

// NewRepository creates  a new repository
func NewRepository() *GitRepository {
	return &GitRepository{}
}
