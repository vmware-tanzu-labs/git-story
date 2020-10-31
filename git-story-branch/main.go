package main

import (
	"fmt"

	storybranch "github.com/carpeliam/git-story-branch"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Printf("This work was done in service of the `%s` story. ", storybranch.GetSpikeStoryName())
	fmt.Printf("It was committed to the `%s` branch.\n", storybranch.GetBranchName())
}
