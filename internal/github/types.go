package github

// RateLimitInfo represents GitHub API rate limit information
type RateLimitInfo struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

// RateLimitResponse represents the response from GitHub's rate limit API
type RateLimitResponse struct {
	Rate RateLimitInfo `json:"rate"`
}

// User represents a GitHub user or organization
type User struct {
	Login       string `json:"login"`
	ID          int64  `json:"id"`
	Type        string `json:"type"` // "User" or "Organization"
	Name        string `json:"name"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   string `json:"created_at"`
}