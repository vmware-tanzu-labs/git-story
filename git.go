package storybranch

import (
	"log"
	"os/exec"
	"strings"
)

// GetBranchName ... gets the branch name. What would we do without this comment.
func GetBranchName() string {
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
		return "" // thar be antipatterns, I'd imagine
	}
	return strings.TrimSpace(string(output))
}
