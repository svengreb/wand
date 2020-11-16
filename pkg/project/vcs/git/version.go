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

const (
	// maxSuitableTagCandidates is the maximum search amount of suitable tag candidates in all commits of the current
	// branch. The value is the same like the default used by Git.
	maxSuitableTagCandidates = 10
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
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to open repository at path %q: %v", repositoryPath, repoOpenErr)
	}

	// Find the latest commit reference of the current branch.
	branchRefs, repoBranchErr := repo.Branches()
	if repoBranchErr != nil {
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to find latest commit reference of the current branch: %v", repoBranchErr)
	}
	headRef, repoHeadErr := repo.Head()
	if repoHeadErr != nil {
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to get the reference where HEAD is pointing to: %v", repoHeadErr)
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
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to iterate over references: %v", branchRefIterErr)
	}

	// Find all commits in the repository starting from HEAD of the current branch.
	commitIterator, commitIterErr := repo.Log(&git.LogOptions{
		From:  currentBranchRef.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if commitIterErr != nil {
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to get the commit history from the current branch: %v", commitIterErr)
	}

	// Query all tags and store them in a temporary map.
	tagIterator, repoTagsErr := repo.Tags()
	if repoTagsErr != nil {
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to get all tag references: %v", repoTagsErr)
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
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to iterate over tags: %v", tagIterErr)
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
			//nolint:errorlint // This is by design to prevent errors from external packages become part of the public
			// API.
			return nil, fmt.Errorf("failed to iterate over commits: %v", tagCommitIterErr)
		}

		if candidate.annotated {
			if tagCandidatesFound < maxSuitableTagCandidates {
				candidate.distance = tagCount
				tagCandidates = append(tagCandidates, candidate)
			}
			tagCandidatesFound++
		}

		if tagCandidatesFound > maxSuitableTagCandidates || len(tags) == 0 {
			break
		}
	}

	// Use the given version by default or...
	semVersion, semVerErr := semver.NewVersion(defaultVersion)
	version := &Version{Version: semVersion}
	if semVerErr != nil {
		//nolint:errorlint // This is by design to prevent errors from external packages become part of the public API.
		return nil, fmt.Errorf("failed to parse default version: %v", semVerErr)
	}
	// ...the latest Git tag from the current branch if at least one tag has been found.
	if len(tagCandidates) != 0 {
		semVersion, semVerErr = semver.NewVersion(tagCandidates[0].ref.Name().Short())
		version = &Version{Version: semVersion}
		if semVerErr != nil {
			//nolint:errorlint // This is by design to prevent errors from external packages become part of the public
			// API.
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
			metadataVersion, mdvErr := version.SetMetadata(fmt.Sprintf("%s-%s", version.Metadata(), buildMetaData))
			if mdvErr != nil {
				//nolint:errorlint // This is by design to prevent errors from external packages become part of the
				// public API.
				return nil, fmt.Errorf("failed to set version metadata: %v", mdvErr)
			}
			version.Version = &metadataVersion
		} else {
			metadataVersion, mdvErr := version.SetMetadata(buildMetaData)
			if mdvErr != nil {
				//nolint:errorlint // This is by design to prevent errors from external packages become part of the
				// public API.
				return nil, fmt.Errorf("failed to set version metadata: %v", mdvErr)
			}
			version.Version = &metadataVersion
		}

		version.CommitsAhead = tagCandidates[0].distance
		version.CommitHash = currentBranchRef.Hash()
		version.LatestVersionTag = tagCandidates[0].ref
	}

	return version, nil
}
