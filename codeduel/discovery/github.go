package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xedom/codeduel/codeduel/utils"
)

type GHPackage struct {
	Name       string `json:"name"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
}

type GitHubProvider struct {
	Org   string
	Repo  string
	Token string
}

func NewGitHubProvider(config *utils.Config) *GitHubProvider {
	return &GitHubProvider{
		Org:   config.GitHubOrg,
		Repo:  config.GitHubRepo,
		Token: config.GitHubToken,
	}
}

func (g *GitHubProvider) GetAvailableLanguages() (map[string]struct{}, error) {
	availableImages := make(map[string]struct{})

	url := fmt.Sprintf("https://api.github.com/orgs/%s/packages?package_type=container", g.Org)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "codeduel-runner-discovery")

	if g.Token != "" {
		req.Header.Set("Authorization", "Bearer "+g.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api returned status %d: %s", resp.StatusCode, string(body))
	}

	var packages []GHPackage
	if err := json.NewDecoder(resp.Body).Decode(&packages); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	for _, pkg := range packages {
		if pkg.Repository.Name == g.Repo {
			availableImages[pkg.Name] = struct{}{}
		}
	}

	return availableImages, nil
}
