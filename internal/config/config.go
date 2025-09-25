package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds the application configuration
type Config struct {
	GitHubToken    string
	MaxConcurrency int
	CloneTimeout   time.Duration
	APITimeout     time.Duration
	RetryAttempts  int
	BaseDir        string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		GitHubToken:    os.Getenv("GITHUB_TOKEN"),
		MaxConcurrency: 5,
		CloneTimeout:   10 * time.Minute,
		APITimeout:     30 * time.Second,
		RetryAttempts:  3,
		BaseDir:        ".",
	}
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	cfg := DefaultConfig()

	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		cfg.GitHubToken = token
	}

	if concurrency := os.Getenv("MAX_CONCURRENCY"); concurrency != "" {
		if val, err := strconv.Atoi(concurrency); err == nil && val > 0 {
			cfg.MaxConcurrency = val
		}
	}

	if timeout := os.Getenv("CLONE_TIMEOUT"); timeout != "" {
		if val, err := time.ParseDuration(timeout); err == nil {
			cfg.CloneTimeout = val
		}
	}

	if baseDir := os.Getenv("BASE_DIR"); baseDir != "" {
		cfg.BaseDir = baseDir
	}

	return cfg
}

// Constants for the application
const (
	AppName    = "mass-git-cloner"
	AppVersion = "1.0.0"
	UserAgent  = AppName + "/" + AppVersion
)

// GitHub API constants
const (
	GitHubAPIBaseURL = "https://api.github.com"
	PerPage          = 100 // Maximum repositories per API request
)
