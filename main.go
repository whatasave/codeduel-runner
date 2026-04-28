package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"github.com/xedom/codeduel/codeduel/api"
	"github.com/xedom/codeduel/codeduel/discovery"
	"github.com/xedom/codeduel/codeduel/runner"
	"github.com/xedom/codeduel/codeduel/utils"
)

func main() {
	loadConfig := utils.LoadConfig()

	provider := &discovery.GitHubProvider{
		Org:   "whatasave",
		Repo:  "codeduel-runner-containers",
		Token: loadConfig.GitHubToken,
	}

	langs, err := provider.GetAvailableLanguages()
	if err != nil {
		log.Printf("[MAIN] Warning: Could not fetch dynamic languages: %v. Falling back to local defaults.", err)
	}

	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("[MAIN] Error initializing Docker client: %v", err)
	}

	codeRunner := runner.NewRunner(dockerCli, langs)

	server, err := api.NewAPIServer(
		loadConfig.Host,
		loadConfig.Port,
		codeRunner,
	)

	if err != nil {
		log.Fatalf("[MAIN] Error creating API server: %v", err)
	}

	server.Run()
}

func loadingEnvVars() {
	isProduction := os.Getenv("GO_ENV") == "production"
	if isProduction {
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("[MAIN] Error getting working directory: %v", err)
	}
	pathDir, err := filepath.Abs(wd)
	if err != nil {
		log.Printf("[MAIN] Error getting absolute path: %v", err)
	}

	log.Printf("Loading .env file from %s", pathDir)
	envPath := filepath.Join(pathDir, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Printf("[MAIN] Error: .env file not found in %s", pathDir)
		return
	}
	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("[MAIN] Error loading .env file")
	}
}

func warnUndefinedEnvVars() {
	envVars := []string{
		"HOST",
		"PORT",
		"DOCKER_IMAGE_PREFIX",
		"DOCKER_TIMEOUT",
	}

	for _, envVar := range envVars {
		test, exists := os.LookupEnv(envVar)
		if !exists {
			log.Printf("[MAIN] Warning: %s not defined in .env file", envVar)
		}
		if test == "" {
			log.Printf("[MAIN] Warning: %s is empty", envVar)
		}
		log.Printf("[MAIN] %s: %s", envVar, test)
	}
}
