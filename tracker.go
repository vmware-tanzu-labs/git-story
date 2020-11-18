package storybranch

import (
	"net/http"

	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

// Tracker comment
type Tracker interface {
	GetStoryDescription(storyID int) string
}

// StoryService comment
type StoryService interface {
	GetByID(storyID int) (*pivotal.Story, *http.Response, error)
}

// PivotalTracker comment
type PivotalTracker struct {
	storyService StoryService
}

// GetStoryDescription comment
func (tracker PivotalTracker) GetStoryDescription(storyID int) string {
	storyService := tracker.storyService
	story, _, _ := storyService.GetByID(storyID)
	return story.Description
}

// NewPivotalTracker returns a new tracker
func NewPivotalTracker(storyService StoryService) *PivotalTracker {
	return &PivotalTracker{storyService}
}
