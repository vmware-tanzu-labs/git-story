# git-story

So you've got a git branch, and it's in support of a particular [Pivotal Tracker](https://www.pivotaltracker.com) user story, so you put your [story ID](https://www.pivotaltracker.com/help/articles/working_with_stories/) at the end of your branch name.

With the power of `git-story`, this allows you to:

1. Open the story for the current branch in your browser
2. Delete old branches for stories that have already been accepted
3. ~~Automatically unsubscribe from unwanted spam emails~~
4. ~~Be reminded if you haven't watered your plants recently~~

At least, it will, some day.

## Installation

1. Configure your [Pivotal Tracker API token](https://www.pivotaltracker.com/help/articles/api_token/)

```sh
export TRACKER_API_TOKEN="YOUR PIVOTAL TRACKER API TOKEN" # https://www.pivotaltracker.com/help/articles/api_token/
```

2. Setup the homebrew tap and install git-story

```sh
$ brew tap git-story-branch/tap
$ brew install git-story
```

## Usage

`git-story` is useful when your git branch name ends in the story ID of the Pivotal Tracker story you're working on:

```sh
$ git checkout WIP-some-story-1234567890
```

### git story-view : view the details of a story

```sh
$ git story-view
State: delivered
Description goes here...
```

### git story-open : open the story URL in your default browser

```sh
$ git story-open
```

## Contributing to Git-Story
We use [ginkgo](https://github.com/onsi/ginkgo) for testing.

To build the CLI tools:

```sh
go build -v -o . ./cli/...
```

To run:

```sh
# assuming the library has been built and you're in a branch that has a story ID at the end...
./git-story-view
./git-story-open
```