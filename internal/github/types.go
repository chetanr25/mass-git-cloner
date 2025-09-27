package github

type User struct {
	Login       string `json:"login"`
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   string `json:"created_at"`
}
