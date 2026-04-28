package main

import (
	_ "embed"
	"log"

	"github.com/docker/docker/client"
	"github.com/xedom/codeduel/codeduel/api"
	"github.com/xedom/codeduel/codeduel/discovery"
	"github.com/xedom/codeduel/codeduel/runner"
	"github.com/xedom/codeduel/codeduel/utils"
)

func main() {
	config := utils.LoadConfig()

	provider := &discovery.GitHubProvider{
		Org:   config.GitHubOrg,
		Repo:  config.GitHubRepo,
		Token: config.GitHubToken,
	}

	langs, err := provider.GetAvailableLanguages()
	if err != nil {
		log.Printf("[MAIN] Warning: Could not fetch dynamic languages: %v. Falling back to local defaults.", err)
	}

	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("[MAIN] Error initializing Docker client: %v", err)
	}

	codeRunner := runner.NewRunner(
		config,
		dockerCli,
		langs,
	)

	server, err := api.NewAPIServer(
		config,
		codeRunner,
	)

	if err != nil {
		log.Fatalf("[MAIN] Error creating API server: %v", err)
	}

	server.Run()
}
