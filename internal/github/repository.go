package github

import "github.com/chetanr25/mass-git-cloner/pkg/models"

// FilterRepositories filters repositories based on the given filter type
func FilterRepositories(repos []*models.Repository, filter models.FilterType) []*models.Repository {
	if filter == models.FilterAll {
		return repos
	}

	filtered := make([]*models.Repository, 0)
	
	for _, repo := range repos {
		switch filter {
		case models.FilterNonForks:
			if !repo.IsFork {
				filtered = append(filtered, repo)
			}
		case models.FilterForksOnly:
			if repo.IsFork {
				filtered = append(filtered, repo)
			}
		}
	}

	return filtered
}

// CalculateStats calculates statistics for a list of repositories
func CalculateStats(repos []*models.Repository) *models.RepositoryStats {
	stats := &models.RepositoryStats{
		Total: len(repos),
	}

	for _, repo := range repos {
		if repo.IsFork {
			stats.Forks++
		} else {
			stats.NonForks++
		}

		if repo.IsPrivate {
			stats.Private++
		} else {
			stats.Public++
		}
	}

	return stats
}

// SortRepositories sorts repositories by various criteria
func SortRepositories(repos []*models.Repository, sortBy string) {
	// Implementation for sorting by name, stars, updated_at, etc.
	// This can be expanded based on requirements
}