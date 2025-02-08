# Codeduel Runner


## Local Setup

Install all the Project dependencies.
```bash
$ go mod download
```

Run the `docker_setup` script inside the root folder. It will check all the langueges present in the `docker/` folder; build a docker image for each of them a save the language name in the `language.txt` file.
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

```
$ docker run ...
```