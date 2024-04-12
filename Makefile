include .$(PWD)/.env
BINARY_NAME=codeduel-runner

build:
	go build -o ./bin/$(BINARY_NAME).exe -v

run: build
	./bin/$(BINARY_NAME)

dev:
	go run .

test:
	go test -v ./...

docker-setup:
	powershell ./docker_setup.ps1

docker-push:
	docker push xedom/codeduel-runner

docker-build:
	docker build --build-arg="PORT=$(PORT)" -t xedom/codeduel-runner .

docker-up: docker-build docker-down
	docker run -d --privileged -v="/var/run/docker.sock:/var/run/docker.sock" -p=$(PORT):$(PORT) --name="codeduel-runner" --env-file=".env" xedom/codeduel-runner

docker-down:
	-docker stop xedom/codeduel-runner && docker rm xedom/codeduel-runner

clean:
	go clean
	-rm -f bin/$(BINARY_NAME).exe
