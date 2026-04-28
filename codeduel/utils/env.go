package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string

	DockerImagePrefix string
	DockerTimeout     string

	GitHubToken string
	GitHubOrg   string
	GitHubRepo  string
}

var config *Config

func LoadConfig() *Config {
	if config == nil {
		// loading env only if not in production
		if GetEnv("ENV", "development") == "development" {
			if err := godotenv.Load(); err != nil {
				log.Printf("%s Error loading .env file", GetLogTag("ENV"))
			}
		}

		config = &Config{
			Host: GetEnv("HOST", "localhost"),
			Port: GetEnv("PORT", "5020"),

			DockerImagePrefix: GetEnv("DOCKER_IMAGE_PREFIX", "cdr-"),
			DockerTimeout:     GetEnv("DOCKER_TIMEOUT", "5"),

			GitHubToken: GetEnv("GITHUB_TOKEN", ""),
			GitHubOrg:   GetEnv("GITHUB_ORG", "whatasave"),
			GitHubRepo:  GetEnv("GITHUB_REPO", "codeduel-runner-containers"),
		}
	}

	return config
}

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("%s%s %s variable not found, using default value: %s\n", GetLogTag("ENV"), GetLogTag("warn"), key, defaultValue)
		return defaultValue
	}

	return value
}
