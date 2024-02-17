BINARY_NAME=codeduel-runner.exe

build:
	go build -o ./bin/$(BINARY_NAME) -v

run: build
	./bin/$(BINARY_NAME)

dev:
	go run .

test:
	go test -v ./...

docker-build:
	docker build -t codeduel-runner .

docker-up:
	docker run -d -p 5001:5001 --name codeduel-runner --env-file .env.docker codeduel-runner

docker-down:
	docker stop codeduel-runner
	docker rm codeduel-runner

docker-restart: docker-down docker-up

clean:
	go clean
	rm -f bin/$(BINARY_NAME)
