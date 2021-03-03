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
	story := usecases.GetStory(adapters.NewRepository(), tracker)
	fmt.Printf("State: %v\n%v", story.State, story.Description)
}
