package storybranch

import (
	"regexp"
	"strconv"
)

func getPivotalTrackerTaskID(branchName string) int {
	re := regexp.MustCompile(`\d+$`)
	taskIDString := re.FindString(branchName)
	taskID, _ := strconv.Atoi(taskIDString)
	return taskID
}

// GetStoryDescription comment
func GetStoryDescription(repo Repository, tracker Tracker) string {
	currentBranchName := repo.GetBranchName()
	storyID := getPivotalTrackerTaskID(currentBranchName)
	return tracker.GetStoryDescription(storyID)
}
