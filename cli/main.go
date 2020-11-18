package main

import (
	"fmt"
	"os"

	storybranch "github.com/carpeliam/git-story-branch"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

func newTracker() *storybranch.PivotalTracker {
	apiToken := os.Getenv("TRACKER_API_TOKEN")
	client := pivotal.NewClient(apiToken)
	return storybranch.NewPivotalTracker(client.Stories)
}

func main() {
	tracker := newTracker()
	description := storybranch.GetStoryDescription(storybranch.NewRepository(), tracker)
	fmt.Printf(description)
}
