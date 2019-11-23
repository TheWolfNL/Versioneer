package cmd

import (
	"strings"

	"github.com/thewolfnl/semver"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type releaseType int

const (
	Major = iota
	Minor
	Patch
)

func (i releaseType) String() string {
	return [...]string{"Major", "Minor", "Patch"}[i]
}

var repo *git.Repository
var latest string

func init() {
	// We instantiate a new repository targeting the given path (the .git folder)
	repository, err := git.PlainOpen("./")
	handleFatal(err)
	// Assign the repo to the global variable
	repo = repository
}

func getLatestVersion() string {
	if latest != "" {
		return latest
	}

	tagrefs, err := repo.Tags()
	handleError(err)
	latestVersion, err := semver.Make("0.0.0")
	handleError(err)
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		version, err := semver.Make(strings.TrimPrefix(strings.TrimPrefix(string(t.Name()), "refs/tags/"), "v"))
		if err != nil {
			return err
		}
		if version.Compare(latestVersion) > 0 {
			latestVersion = version
		}
		return nil
	})
	handleError(err)
	latest = latestVersion.String()
	return latest
}

func incrementVersion(releaseType releaseType) string {
	prevVersion := latest
	newVersion, err := semver.Make(prevVersion)
	handleError(err)

	output("Increment Version")

	switch releaseType {
	case Major:
		newVersion.IncrementMajor()
	case Minor:
		newVersion.IncrementMinor()
	case Patch:
		newVersion.IncrementPatch()
	}
	return newVersion.String()
}

func createAndPushTag(tagName string) {
	head, err := repo.Head()
	handleFatal(err)

	output("Tagging " + tagName)
	_, err = repo.CreateTag(tagName, head.Hash(), nil)
	handleFatal(err)

	output("Pushing tag " + tagName)
	err = repo.Push(&git.PushOptions{
		RefSpecs: []config.RefSpec{"refs/tags/*:refs/tags/*"},
	})
	handleFatal(err)
}
