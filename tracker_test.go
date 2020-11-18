package storybranch_test

import (
	"errors"
	"net/http"

	storybranch "github.com/carpeliam/git-story-branch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

type PivotalTrackerStoryServiceStub struct{}

func (pivotalTrackerStoryServiceStub PivotalTrackerStoryServiceStub) GetByID(storyID int) (*pivotal.Story, *http.Response, error) {
	if storyID == 123456789 {
		newStory := pivotal.Story{
			ID:            123456789,
			ProjectID:     0,
			Name:          "My cool story",
			Description:   "I dunno, uh, cool story... bro.. or something.",
			Type:          "Type",
			State:         "State",
			Estimate:      nil,
			AcceptedAt:    nil,
			Deadline:      nil,
			RequestedByID: 0,
			OwnerIDs:      []int{},
			LabelIDs:      []int{},
			Labels:        []*pivotal.Label{},
			TaskIDs:       []int{},
			Tasks:         []int{},
			FollowerIDs:   []int{},
			CommentIDs:    []int{},
			CreatedAt:     nil,
			UpdatedAt:     nil,
			BeforeID:      0,
			AfterID:       0,
			IntegrationID: 0,
			ExternalID:    "ExternalID",
			URL:           "URL",
		}
		return &newStory, nil, nil
	}

	return nil, nil, errors.New("Story does not exist.")
}

var _ = Describe("Tracker", func() {

	It("should be able to look up the description of a story given the story ID", func() {
		tracker := storybranch.NewPivotalTracker(PivotalTrackerStoryServiceStub{})

		description := tracker.GetStoryDescription(123456789)
		Expect(description).To(Equal("I dunno, uh, cool story... bro.. or something."))
	})
})
