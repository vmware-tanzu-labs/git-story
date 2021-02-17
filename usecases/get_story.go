package usecases

import (
	"regexp"
	"strconv"
)

type Story struct {
	ID int
	Description string
	State string
}

func getPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, _ := strconv.Atoi(taskIDString)
	return taskID
}

// GetStory comment
func GetStory(repo Repository, tracker Tracker) *Story {
	currentBranchName := repo.GetBranchName()
	storyID := getPivotalTrackerTaskID(currentBranchName)
	
	return tracker.GetStory(storyID)
}

// Tracker comment
type Tracker interface {
	GetStory(storyID int) *Story
}

// Repository comment
type Repository interface {
	GetBranchName() string
}