package adapters

import (
	"net/http"

	"github.com/carpeliam/git-story-branch/usecases"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

// StoryService comment
type StoryService interface {
	GetByID(storyID int) (*pivotal.Story, *http.Response, error)
}

// PivotalTracker comment
type PivotalTracker struct {
	storyService StoryService
}

// GetStory comment
func (tracker PivotalTracker) GetStory(storyID int) *usecases.Story {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	
	return &usecases.Story{
		ID: story.ID,
		Description: story.Description,
		State: story.State,
	}
}

// NewPivotalTracker returns a new tracker
func NewPivotalTracker(storyService StoryService) usecases.Tracker {
	return &PivotalTracker{storyService}
}
