package adapters

import (
	"errors"
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

func (repo GitRepository) DeleteBranch(branchName string) error {
	cmd := exec.Command("git", "branch", "-D", branchName)
	output, error := cmd.CombinedOutput()
	if error != nil {
		return errors.New(strings.Trim(string(output), "\n"))
	}
	return nil
}
func (repo GitRepository) GetAllBranchNames() []string {
	output, _ := exec.Command("git", "--no-pager", "branch", "--format", "%(refname:short)").Output()
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

// NewRepository creates  a new repository
func NewRepository() *GitRepository {
	return &GitRepository{}
}
