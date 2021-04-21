package usecases

import "fmt"

func OpenStory(repo Repository, tracker Tracker, browser Browser) error {
	story, error := GetStory(repo, tracker)
	if error != nil {
		return error
	}
	fmt.Println("URL", story.URL)
	// exec.Command("open", story.URL).Start()
	return browser.OpenURL(story.URL)
}

type Browser interface {
	OpenURL(URL string) error
}
