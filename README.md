# Codeduel Runner

## Local Setup

Install all the Project dependencies.

```bash
$ go mod download
```

Now you can run the Project.

```bash
$ go run .
```

## Docker Setup

```bash
$ docker build -t xedom/codeduel-runner .

$ docker run -p 5020:80 --env-file .env.docker -v /var/run/docker.sock:/var/run/docker.sock xedom/codeduel-runner
```
