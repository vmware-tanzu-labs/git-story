# git-story-branch

So you've got a branch, and it's in support of a particular user story, so you put your story id at the end of your branch name.

With the power of `git-story-branch`, this allows you to:

1. Open the story for the current branch in your browser
2. Delete old branches for stories that have already been accepted
3. ~~Automatically unsubscribe from unwanted spam emails~~
4. ~~Be reminded if you haven't watered your plants recently~~

At least, it will, some day.

## Development

### Workspace
Go is opinionated about where you put your stuff, so here's where I put my stuff. I don't really know what I'm doing, maybe you can help me make it better or confirm that this isn't that bad after all.
```sh
export GOPATH=~/workspace/go # or wherever you feel like putting your stuff
mkdir -p ${GOPATH}/src/github.com/carpeliam
git clone https://github.com/carpeliam/git-story-branch.git ${GOPATH}/src/github.com/carpeliam/git-story-branch
```

### Running tests
This uses [ginkgo](http://onsi.github.io/ginkgo/). So, you can just run `ginkgo`.

### Running the app
For now:
```sh
go run ./git-story-branch/main.go
```
