package utils

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullImages(config *Config, dockerCli *client.Client, images map[string]string) {
	if dockerCli == nil {
		log.Printf("[UTILS] Docker client is not initialized")
		return
	}

	var wg sync.WaitGroup
	for lang, imageName := range images {
		wg.Add(1)
		go func(lang, imageName string) {
			defer wg.Done()
			log.Printf("[UTILS] Pulling image for %s: %s", lang, imageName)
			reader, err := dockerCli.ImagePull(context.Background(), imageName, image.PullOptions{})
			if err != nil {
				log.Printf("[UTILS] Error pulling image for %s: %v", lang, err)
				return
			}

			_, _ = io.Copy(io.Discard, reader)
			_ = reader.Close()
			log.Printf("[UTILS] Successfully pulled image for %s", lang)
		}(lang, imageName)
	}

	wg.Wait()
}
