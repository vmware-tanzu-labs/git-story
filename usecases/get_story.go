package usecases

import (
	"errors"
	"regexp"
	"strconv"
)

func getPivotalTrackerTaskID(branchName string) (int, error) {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, error := strconv.Atoi(taskIDString)
	return taskID, error
}

// GetStory comment
func GetStory(repo Repository, tracker Tracker) (*Story, error) {
	currentBranchName := repo.GetCurrentBranchName()
	return GetStoryByBranchName(currentBranchName, tracker)
}

func GetStoryByBranchName(branchName string, tracker Tracker) (*Story, error) {
	storyID, branchError := getPivotalTrackerTaskID(branchName)

	if branchError != nil {
		return nil, errors.New("please run in branch that contains a Pivotal Tracker Story ID")
	}

	return tracker.GetStory(storyID)
}
