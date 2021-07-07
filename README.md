# git-story

So you've got a git branch, and it's in support of a particular [Pivotal Tracker](https://www.pivotaltracker.com) user story, so you put your [story ID](https://www.pivotaltracker.com/help/articles/working_with_stories/) at the end of your branch name.

With the power of `git-story`, this allows you to:

1. Open the story for the current branch in your browser
2. Delete old branches for stories that have already been accepted
3. ~~Automatically unsubscribe from unwanted spam emails~~
4. ~~Be reminded if you haven't watered your plants recently~~

## What? Why?

Git branches. We've all got 'em. And when you've got a lot of them, it can get kind of confusing:

```sh
git branch
#   EOD-sockjs-spike
#   WIP-service-refactor
#   feature/ui-rewrite
#   feature/ui-rewrite-fixed
# * main
#   uiRewrite
```
Which branch should we use for that UI Rewrite feature? And how do we know if any of these branches can be deleted?

I humbly suggest that you **add your story ID to your branch name**, so that it becomes easier to identify which story you're actively working on, and easier to identify whether or not a branch corresponds to an accepted feature and can be safely deleted. If `feature/ui-rewrite` was instead `feature/ui-rewrite-1234567890`, we'd know exactly which story it was associated with.

If you take me up on this suggestion, then `git-story` becomes a truly useful tool!

## Installation

1. Configure your [Pivotal Tracker API token](https://www.pivotaltracker.com/help/articles/api_token/). This is necessary if your Tracker project is private.

```sh
export TRACKER_API_TOKEN="YOUR PIVOTAL TRACKER API TOKEN" # https://www.pivotaltracker.com/help/articles/api_token/
```

2. Install `git-story` via Homebrew

```sh
brew install git-story-branch/tap/git-story
```

## Usage

Assuming that the name of your current branch ends in a story ID, you can do all of the following cool tricks from the comfort of your terminal:

### git story-view : view the details of a story

```sh
git story-view
# Example output:
# State: delivered
# Description goes here...
```

### git story-open : open the story URL in your default browser

```sh
git story-open
```

### git story-sweep : delete all local branches corresponding to accepted stories, and report any errors in deleting

```sh
git story-sweep
# Example output:
# Could not delete 'WIP-accepted-branch-#987654321', error: Cannot delete branch 'WIP-accepted-branch-#987654321' checked out at '/path'  # This is an expected error when you're deleting the current branch
# Deleted 'accepted-feature-branch-#123456789'
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
