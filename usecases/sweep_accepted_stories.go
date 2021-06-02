package usecases

// import "log"

// get all the branch names

// call get story on each of those things

// see if they're accepted

// delete them if they are

func SweepAcceptedStories(repo Repository, tracker Tracker) error {
	branchNames := repo.GetAllBranchNames()

	for _, branchName := range branchNames {
		story, _ := GetStoryByBranchName(branchName, tracker)

		if story != nil && story.State == "accepted" {
			_, error := repo.DeleteBranch(branchName)

			if error != nil {
				return error
			}
		}
	}

	return nil
}
