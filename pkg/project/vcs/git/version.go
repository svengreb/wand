// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package git

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

// Version stores version information and metadata derived from a Git repository.
type Version struct {
	// Version is the semantic version.
	// See https://semver.org for more details.
	*semver.Version

	// CommitsAhead is the count of commits ahead to the latest Git version tag in the current branch.
	CommitsAhead int

	// CommitHash is the hash of the latest commit in the current branch.
	CommitHash plumbing.Hash

	// LatestVersionTag is the latest Git version tag in the current branch.
	LatestVersionTag *plumbing.Reference
}

// deriveVersion derives version information and metadata from a Git repository.
// It searches for the latest SemVer (https://semver.org) compatible version tag in the current branch and falls
// back to the given default version if no tag is found.
// If at least one tag is found, but it is not the latest commit of the current branch, the build metadata is appended,
// consisting of the amount of commits ahead and the shortened reference hash (8 digits) of the latest commit from the
// current branch.
//
// This function is an early implementation of the Git `describe` command because support in `go-git` has not been
// implemented yet. See the full compatibility comparison documentation with Git at
// https://github.com/go-git/go-git/blob/master/COMPATIBILITY.md as well as the proposed Git `describe` command
// implementation at https://github.com/src-d/go-git/pull/816 for more details.
func deriveVersion(defaultVersion, repositoryPath string) (*Version, error) {
	if defaultVersion == "" {
		return nil, fmt.Errorf("default version must not be empty")
	}

	repo, repoOpenErr := git.PlainOpen(repositoryPath)
	if repoOpenErr != nil {
		return nil, repoOpenErr
	}

	// Find the latest commit reference of the current branch.
	branchRefs, repoBranchErr := repo.Branches()
	if repoBranchErr != nil {
		return nil, repoBranchErr
	}
	headRef, repoHeadErr := repo.Head()
	if repoHeadErr != nil {
		return nil, repoHeadErr
	}
	var currentBranchRef plumbing.Reference
	branchRefIterErr := branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchRef = *branchRef
			return nil
		}
		return nil
	})
	if branchRefIterErr != nil {
		return nil, branchRefIterErr
	}

	// Find all commits in the repository starting from HEAD of the current branch.
	commitIterator, commitIterErr := repo.Log(&git.LogOptions{
		From:  currentBranchRef.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if commitIterErr != nil {
		return nil, commitIterErr
	}

	// Query all tags and store them in a temporary map.
	tagIterator, repoTagsErr := repo.Tags()
	if repoTagsErr != nil {
		return nil, repoTagsErr
	}
	tags := make(map[plumbing.Hash]*plumbing.Reference)
	tagIterErr := tagIterator.ForEach(func(tag *plumbing.Reference) error {
		if tagObject, tagObjectErr := repo.TagObject(tag.Hash()); tagObjectErr == nil {
			// Only include tags that have a valid SemVer version format.
			if _, semVerParseErr := semver.NewVersion(tag.Name().Short()); semVerParseErr == nil {
				tags[tagObject.Target] = tag
			}
		} else {
			tags[tag.Hash()] = tag
		}
		return nil
	})
	tagIterator.Close()
	if tagIterErr != nil {
		return nil, tagIterErr
	}

	type describeCandidate struct {
		ref       *plumbing.Reference
		annotated bool
		distance  int
	}
	var tagCandidates []*describeCandidate
	var tagCandidatesFound int
	tagCount := -1

	// Search for maximal 10 (Git default) suitable tag candidates in all commits of the current branch.
	for {
		candidate := &describeCandidate{annotated: false}
		tagCommitIterErr := commitIterator.ForEach(func(commit *object.Commit) error {
			tagCount++
			if tagReference, ok := tags[commit.Hash]; ok {
				delete(tags, commit.Hash)
				candidate.ref = tagReference
				hash := tagReference.Hash()
				if !bytes.Equal(commit.Hash[:], hash[:]) {
					candidate.annotated = true
				}
				return storer.ErrStop
			}
			return nil
		})
		if tagCommitIterErr != nil {
			return nil, tagCommitIterErr
		}

		if candidate.annotated {
			if tagCandidatesFound < 10 {
				candidate.distance = tagCount
				tagCandidates = append(tagCandidates, candidate)
			}
			tagCandidatesFound++
		}

		if tagCandidatesFound > 10 || len(tags) == 0 {
			break
		}
	}

	// Use the given version by default or...
	semVersion, semVerErr := semver.NewVersion(defaultVersion)
	version := &Version{Version: semVersion}
	if semVerErr != nil {
		return nil, fmt.Errorf("failed to parse default version: %v", semVerErr)
	}
	// ...the latest Git tag from the current branch if at least one tag has been found.
	if len(tagCandidates) != 0 {
		semVersion, semVerErr = semver.NewVersion(tagCandidates[0].ref.Name().Short())
		version = &Version{Version: semVersion}
		if semVerErr != nil {
			return nil, fmt.Errorf("failed to parse version from Git tag %s: %v",
				tagCandidates[0].ref.Name().Short(), semVerErr)
		}
	}
	// Add additional version information if the latest commit of the current branch is not the found tag.
	if len(tagCandidates) != 0 && tagCandidates[0].distance > 0 {
		// If not included in the tag already, append metadata consisting of the amount of commit(s) ahead and the
		// shortened commit hash (8 digits) of the latest commit.
		buildMetaData := fmt.Sprintf("%s.%s",
			strconv.Itoa(tagCandidates[0].distance), currentBranchRef.Hash().String()[:8])
		if version.Metadata() != "" {
			metadataVersion, err := version.SetMetadata(fmt.Sprintf("%s-%s", version.Metadata(), buildMetaData))
			if err != nil {
				return nil, err
			}
			version.Version = &metadataVersion
		} else {
			metadataVersion, err := version.SetMetadata(buildMetaData)
			if err != nil {
				return nil, err
			}
			version.Version = &metadataVersion
		}

		version.CommitsAhead = tagCandidates[0].distance
		version.CommitHash = currentBranchRef.Hash()
		version.LatestVersionTag = tagCandidates[0].ref
	}

	return version, nil
}
