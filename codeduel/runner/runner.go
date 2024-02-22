package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Runner struct {
	client *client.Client
	images map[string]struct{}
}

type ExecutionResult struct {
	Output     string `json:"output"`
	Error      string `json:"error"`
	Terminated bool   `json:"terminated"`
}

func NewRunner() (*Runner, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	images := make(map[string]struct{}, 0)
	files, err := os.ReadDir("./docker")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), "_") {
			images[file.Name()] = struct{}{}
		}
	}
	return &Runner{client, images}, nil
}

func (r *Runner) Run(language string, code string, input []string) (*[]ExecutionResult, error) {
	_, ok := r.images[language]
	if !ok {
		return nil, fmt.Errorf("language %s not supported", language)
	}
	runnerContainer, err := r.client.ContainerCreate(context.Background(), &container.Config{
		Image: os.Getenv("DOCKER_IMAGE_PREFIX") + language,
		Env: []string{
			fmt.Sprintf("CODE=%s", code),
			fmt.Sprintf("INPUT=%s", encodeInput(input)),
			fmt.Sprintf("TIMEOUT=%s", os.Getenv("DOCKER_TIMEOUT")),
		},
	}, nil, nil, nil, "")
	if err != nil {
		return nil, err
	}
	if err := r.client.ContainerStart(context.Background(), runnerContainer.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}
	reader, err := r.client.ContainerLogs(context.Background(), runnerContainer.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	outputBuffer := new(strings.Builder)
	errorBuffer := new(strings.Builder)
	_, err = stdcopy.StdCopy(outputBuffer, errorBuffer, reader)
	if err != nil {
		return nil, err
	}
	error := errorBuffer.String()
	output := outputBuffer.String()
	var result []ExecutionResult
	err = json.Unmarshal([]byte(output), &result)
	if err := r.client.ContainerStop(context.Background(), runnerContainer.ID, container.StopOptions{}); err != nil {
		return &result, err
	}
	if err := r.client.ContainerRemove(context.Background(), runnerContainer.ID, types.ContainerRemoveOptions{}); err != nil {
		return &result, err
	}
	if error != "" {
		return nil, fmt.Errorf("container error: %s", error)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot parse result %s: %s", output, err)
	}
	return &result, nil
}

func encodeInput(inputs []string) string {
	var sb strings.Builder
	for i, input := range inputs {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(strings.ReplaceAll(input, "\n", "\\\n"))
	}
	return sb.String()
}
