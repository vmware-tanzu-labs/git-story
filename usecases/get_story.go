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
	URL         string
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
	return GetStoryByBranchName(currentBranchName, tracker)
}

func GetStoryByBranchName(branchName string, tracker Tracker) (*Story, error) {
	storyID, branchError := getPivotalTrackerTaskID(branchName)

	if branchError != nil {
		return nil, errors.New("please run in branch that contains a Pivotal Tracker Story ID")
	}

	return tracker.GetStory(storyID)
}

// Tracker comment
type Tracker interface {
	GetStory(storyID int) (*Story, error)
}

// Repository comment TODO maybe move this to a different file?
type Repository interface {
	GetBranchName() string
	DeleteBranch(branchName string)
	GetAllBranchNames() []string
}
