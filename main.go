package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/xedom/codeduel/codeduel/api"
)

func main() {
	loadingEnvVars()
	warnUndefinedEnvVars()

	server, err := api.NewAPIServer(os.Getenv("HOST"), os.Getenv("PORT"))
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
