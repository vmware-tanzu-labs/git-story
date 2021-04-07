package adapters

import (
	"net/http"

	"github.com/git-story-branch/git-story-branch/usecases"
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
func (tracker PivotalTracker) GetStory(storyID int) (*usecases.Story, error) {
	storyService := tracker.storyService
	story, _, error := storyService.GetByID(storyID)

	if error != nil {
		return nil, error
	}
	return &usecases.Story{
		ID:          story.ID,
		Description: story.Description,
		State:       story.State,
	}, nil
}

// NewPivotalTracker returns a new tracker
func NewPivotalTracker(storyService StoryService) usecases.Tracker {
	return &PivotalTracker{storyService}
}
