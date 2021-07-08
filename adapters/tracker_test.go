package adapters_test

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu-labs/git-story/adapters"
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
			State:         "delivered",
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
			URL:           "https://best-url-ever.com",
		}
		return &newStory, nil, nil
	}

	return nil, nil, errors.New("Story does not exist.")
}

var _ = Describe("Tracker", func() {

	It("should be able to return a story given an ID", func() {
		tracker := adapters.NewPivotalTracker(PivotalTrackerStoryServiceStub{})
		storyID := 123456789

		story, error := tracker.GetStory(storyID)
		Expect(error).To(BeNil())
		Expect(story.Description).To(Equal("I dunno, uh, cool story... bro.. or something."))
		Expect(story.State).To(Equal("delivered"))
		Expect(story.ID).To(Equal(storyID))
		Expect(story.URL).To(Equal("https://best-url-ever.com"))
	})

	It("should return an error if no story can be found for that ID", func() {
		tracker := adapters.NewPivotalTracker(PivotalTrackerStoryServiceStub{})
		storyID := 8675309

		story, error := tracker.GetStory(storyID)
		Expect(story).To(BeNil())
		Expect(error).NotTo(BeNil())
	})
})
