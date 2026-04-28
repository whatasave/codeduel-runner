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

	provider := discovery.NewGitHubProvider(config)

	images, err := provider.GetAvailableImages()
	if err != nil {
		log.Printf("[MAIN] Warning: Could not fetch dynamic images: %v. Falling back to local defaults.", err)
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("[MAIN] Error initializing Docker client: %v", err)
	}

	utils.PullImages(config, dockerClient, images)

	codeRunner := runner.NewRunner(
		config,
		dockerClient,
		images,
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
