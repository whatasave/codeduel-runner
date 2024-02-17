package runner

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Runner struct {
	client     *client.Client
	containers map[string]container.CreateResponse
}

type ExecutionResult struct {
	output string
	error  string
}

var languages = []string{"python", "javascript"}

func NewRunner() (*Runner, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	containers := make(map[string]container.CreateResponse)
	for _, language := range languages {
		container, err := client.ContainerCreate(context.Background(), &container.Config{
			Image: language,
		}, nil, nil, nil, "")
		if err != nil {
			return nil, err
		}
		containers[language] = container
	}

	return &Runner{client, containers}, nil
}

func (r *Runner) Run(language string, code string, input string) (*ExecutionResult, error) {
	error := ""
	output := "test"
	return &ExecutionResult{output, error}, nil
}
