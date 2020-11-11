package storybranch

import (
	"os/exec"
	"strings"
)

// GetBranchName gets the current branch name of the repository.
func GetBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(output))
}
