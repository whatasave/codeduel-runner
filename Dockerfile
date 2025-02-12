FROM golang:1.22 AS build-stage

ENV BINARY_NAME=codeduel-runner
ENV ENV=production

RUN useradd -u 1001 -m codeduel-user

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/$BINARY_NAME -v


FROM build-stage AS run-test-stage
RUN go test -v ./...


FROM docker:20.10.24-dind AS release-stage

WORKDIR /app

COPY --from=build-stage /usr/src/app/bin /usr/local/bin
COPY docker docker
COPY docker_setup.sh docker_setup.sh

ENV BINARY_NAME="codeduel-runner"
ENV DOCKER_IMAGE_PREFIX="cdr-"
ENV DOCKER_TIMEOUT="5s"
ENV ENV="production"
ENV HOST=0.0.0.0
ENV PORT=80

COPY --from=build-stage /etc/passwd /etc/passwd

COPY docker docker
COPY docker_setup.sh docker_setup.sh

RUN chmod +x docker_setup.sh

EXPOSE $PORT

ENTRYPOINT ["sh", "-c", "./docker_setup.sh && codeduel-runner"]
