# Codeduel Runner

## Local Setup

Install all the Project dependencies.

```bash
$ go mod download
```

Run the `docker_setup` script inside the root folder. It will check all the languages present in the `docker/` folder; build a docker image for each of them a save the language name in the `language.txt` file.

```bash
# on windows:
$ .\docker_setup.ps1

# on linux:
$ ./docker_setup.sh
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
