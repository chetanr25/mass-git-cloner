package models

import "time"

// Repository represents a GitHub repository
type Repository struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	FullName      string    `json:"full_name"`
	Description   string    `json:"description"`
	CloneURL      string    `json:"clone_url"`
	SSHURL        string    `json:"ssh_url"`
	Language      string    `json:"language"`
	StarCount     int       `json:"stargazers_count"`
	ForkCount     int       `json:"forks_count"`
	IsFork        bool      `json:"fork"`
	IsPrivate     bool      `json:"private"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Size          int       `json:"size"`
	DefaultBranch string    `json:"default_branch"`
	Selected      bool      `json:"-"`
}

type RepositoryStats struct {
	Total    int
	Forks    int
	NonForks int
	Private  int
	Public   int
}

type FilterType int

const (
	FilterAll FilterType = iota
	FilterNonForks
	FilterForksOnly
)

func (f FilterType) String() string {
	switch f {
	case FilterAll:
		return "All repositories"
	case FilterNonForks:
		return "Non-fork repositories only"
	case FilterForksOnly:
		return "Fork repositories only"
	default:
		return "Unknown"
	}
}

type CloneResult struct {
	Repository *Repository
	Success    bool
	Error      error
	Duration   time.Duration
}

type CloneProgress struct {
	Total     int
	Completed int
	Failed    int
	Current   string
}
