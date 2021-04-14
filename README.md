# git-story-branch

So you've got a branch, and it's in support of a particular user story, so you put your story id at the end of your branch name.

With the power of `git-story-branch`, this allows you to:

1. Open the story for the current branch in your browser
2. Delete old branches for stories that have already been accepted
3. ~~Automatically unsubscribe from unwanted spam emails~~
4. ~~Be reminded if you haven't watered your plants recently~~

At least, it will, some day.

## Installation

1. Configure your tracker API token

```sh
export TRACKER_API_TOKEN="YOUR PIVOTAL TRACKER API TOKEN" # https://www.pivotaltracker.com/help/articles/api_token/
```

2. Setup the homebrew tap and install Git-Story

```sh
$ brew tap git-story-branch/tap
$ brew install git-story
```

3. Use the git subcommand to get the story!


```sh
git story-view
```

## Contributing to Git-Story

To build the source files

```sh
go build ./cli/main.go
```

To run

```sh
./main # assuming the library has been built and you're in a branch that has a story ID at the end
```