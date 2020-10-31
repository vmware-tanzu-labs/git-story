package storybranch

import (
	"log"
	"os"

	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

// GetSpikeStoryName gets the name of the spike. Documentation!
func GetSpikeStoryName() string {
	// given that this spike is looking something up from a private project, we don't even need an api token
	var apiToken = os.Getenv("TRACKER_API_TOKEN")
	var client = pivotal.NewClient(apiToken)
	var story, _, err = client.Stories.Get(2474798, 175526301)
	if err != nil {
		log.Fatal(err)
		return "" // really? is this the best you can do? TODO learn some idioms
	}
	return story.Name
}
