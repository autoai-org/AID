package vcs

import "time"

// GitCommit is a single object for a commit
type GitCommit struct {
	Author  string    `json:"author"`
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}

// GitCommits is an array of GitCommit object
type GitCommits []GitCommit
