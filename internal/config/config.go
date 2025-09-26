package config

import (
	"time"
)

type Config struct {
	CloneTimeout time.Duration
	APITimeout   time.Duration
	BaseDir      string
}

func DefaultConfig() *Config {
	return &Config{
		CloneTimeout: 10 * time.Minute,
		APITimeout:   30 * time.Second,
		BaseDir:      ".",
	}
}

const (
	AppName          = "mass-git-cloner"
	AppVersion       = "1.0.0"
	UserAgent        = AppName + "/" + AppVersion
	GitHubAPIBaseURL = "https://api.github.com"
	PerPage          = 100
)
