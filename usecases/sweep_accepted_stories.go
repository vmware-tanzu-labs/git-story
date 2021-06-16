package usecases

type ResultMap map[string]error

// SweepAcceptedStories returns a map of the branches it attempted to delete and an error if that branch was unable to be deleted
func SweepAcceptedStories(repo Repository, tracker Tracker) ResultMap {
	branchErrors := make(map[string]error)

	branchNames := repo.GetAllBranchNames()

	for _, branchName := range branchNames {
		story, _ := GetStoryByBranchName(branchName, tracker)

		if story != nil && story.State == "accepted" {
			branchErrors[branchName] = repo.DeleteBranch(branchName)
		}
	}

	return branchErrors
}
