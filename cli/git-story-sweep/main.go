package main

import (
	"fmt"
	"os"

	"github.com/git-story-branch/git-story-branch/adapters"
	"github.com/git-story-branch/git-story-branch/usecases"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

func newTracker() usecases.Tracker {
	apiToken := os.Getenv("TRACKER_API_TOKEN")
	client := pivotal.NewClient(apiToken)
	return adapters.NewPivotalTracker(client.Stories)
}

func main() {
	tracker := newTracker()
	results := usecases.SweepAcceptedStories(adapters.NewRepository(), tracker)

	hasError := false
	for branchName, error := range results {
		if error != nil {
			hasError = true
			fmt.Printf("Could not delete '%s', %s\n", branchName, error)
		} else {
			fmt.Printf("Deleted '%s'\n", branchName)
		}
	}
	if hasError {
		os.Exit(1)
	}
}
