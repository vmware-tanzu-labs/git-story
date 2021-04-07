package usecases

import (
	"errors"
	"regexp"
	"strconv"
)

type Story struct {
	ID          int
	Description string
	State       string
}

func getPivotalTrackerTaskID(branchName string) (int, error) {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, error := strconv.Atoi(taskIDString)
	return taskID, error
}

// GetStory comment
func GetStory(repo Repository, tracker Tracker) (*Story, error) {
	currentBranchName := repo.GetBranchName()
	storyID, branchError := getPivotalTrackerTaskID(currentBranchName)

	if branchError != nil {
		return nil, errors.New("Please run in branch that contains a Pivotal Tracker Story ID")
	}

	return tracker.GetStory(storyID)
}

// Tracker comment
type Tracker interface {
	GetStory(storyID int) (*Story, error)
}

// Repository comment
type Repository interface {
	GetBranchName() string
}
