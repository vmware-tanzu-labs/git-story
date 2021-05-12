package usecases

type Story struct {
	ID          int
	Description string
	State       string
	URL         string
}

type Tracker interface {
	GetStory(storyID int) (*Story, error)
}

type Repository interface {
	GetCurrentBranchName() string
	DeleteBranch(branchName string) error
	GetAllBranchNames() []string
}
