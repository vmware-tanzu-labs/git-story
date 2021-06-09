package usecases

func SweepAcceptedStories(repo Repository, tracker Tracker) error {
	branchNames := repo.GetAllBranchNames()

	for _, branchName := range branchNames {
		story, _ := GetStoryByBranchName(branchName, tracker)

		if story != nil && story.State == "accepted" {
			error := repo.DeleteBranch(branchName)
			if error != nil {
				return error
			}
		}
	}

	return nil
}
