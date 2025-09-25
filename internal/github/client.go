package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// Client represents a GitHub API client
type Client struct {
	httpClient *http.Client
	token      string
	baseURL    string
}

// NewClient creates a new GitHub API client
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: cfg.APITimeout,
		},
		token:   cfg.GitHubToken,
		baseURL: config.GitHubAPIBaseURL,
	}
}

// UserExists checks if a GitHub user or organization exists
func (c *Client) UserExists(username string) (bool, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// GetRepositories fetches all repositories for a user or organization
func (c *Client) GetRepositories(username string) ([]*models.Repository, error) {
	var allRepos []*models.Repository
	page := 1

	for {
		repos, hasMore, err := c.getRepositoriesPage(username, page)
		if err != nil {
			return nil, err
		}

		allRepos = append(allRepos, repos...)

		if !hasMore {
			break
		}
		page++
	}

	return allRepos, nil
}

// getRepositoriesPage fetches a single page of repositories
func (c *Client) getRepositoriesPage(username string, page int) ([]*models.Repository, bool, error) {
	url := fmt.Sprintf("%s/users/%s/repos?per_page=%d&page=%d&sort=updated",
		c.baseURL, username, config.PerPage, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, false, err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("GitHub API error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	var repos []*models.Repository
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, false, err
	}

	// Check if there are more pages
	hasMore := len(repos) == config.PerPage

	return repos, hasMore, nil
}

// GetRateLimit returns the current rate limit status
func (c *Client) GetRateLimit() (*RateLimitInfo, error) {
	url := fmt.Sprintf("%s/rate_limit", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rate limit API error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rateLimitResp RateLimitResponse
	if err := json.Unmarshal(body, &rateLimitResp); err != nil {
		return nil, err
	}

	return &rateLimitResp.Rate, nil
}

// setHeaders sets the required headers for GitHub API requests
func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", config.UserAgent)

	if c.token != "" {
		req.Header.Set("Authorization", "token "+c.token)
	}
}

// WaitForRateLimit waits if rate limit is exceeded
func (c *Client) WaitForRateLimit() error {
	rateLimit, err := c.GetRateLimit()
	if err != nil {
		return err
	}

	if rateLimit.Remaining == 0 {
		waitTime := time.Until(time.Unix(rateLimit.Reset, 0))
		if waitTime > 0 {
			time.Sleep(waitTime + time.Second) // Add 1 second buffer
		}
	}

	return nil
}
