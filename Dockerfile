ARG PORT=5000

FROM golang:1.21 as build-stage

WORKDIR /app

COPY . .
RUN go build -o ./bin/codeduel-runner -v

FROM build-stage AS run-test-stage

RUN go test -v ./...

FROM debian AS release-stage

WORKDIR /app

RUN apt-get update && apt-get install -y docker.io

COPY --from=build-stage /app/bin /usr/local/bin
COPY docker docker
COPY docker_setup.sh docker_setup.sh
COPY .env .env

ENV BINARY_NAME="codeduel-runner"
ENV DOCKER_IMAGE_PREFIX="cd-runner-"
ENV DOCKER_TIMEOUT="5s"
ENV ENV="production"
ENV HOST=0.0.0.0
ENV PORT=80

EXPOSE $PORT

ENTRYPOINT ["bash", "-c", "./docker_setup.sh && codeduel-runner"]
