package utils

import (
	"context"
	"io"
	"log"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullImages(config *Config, dockerCli *client.Client, images map[string]string) {
	if dockerCli == nil {
		log.Printf("[UTILS] Docker client is not initialized")
		return
	}

	for lang, imageName := range images {
		log.Printf("[UTILS] Pulling image for %s: %s", lang, imageName)
		reader, err := dockerCli.ImagePull(context.Background(), imageName, image.PullOptions{})
		if err != nil {
			log.Printf("[UTILS] Error pulling image for %s: %v", lang, err)
			continue
		}
		// Consume the stream so the pull actually completes.
		_, _ = io.Copy(io.Discard, reader)
		_ = reader.Close()
		log.Printf("[UTILS] Successfully pulled image for %s", lang)
	}
}
