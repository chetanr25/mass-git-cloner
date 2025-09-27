package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

type Client struct {
	httpClient *http.Client

	baseURL string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: cfg.APITimeout,
		},

		baseURL: config.GitHubAPIBaseURL,
	}
}

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

	hasMore := len(repos) == config.PerPage

	return repos, hasMore, nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", config.UserAgent)
}
