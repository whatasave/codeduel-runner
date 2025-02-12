#!/bin/sh

touch languages.txt

CURRENT_DIR=$(pwd)
find docker -mindepth 1 -maxdepth 1 -type d ! -name "_base" | while read -r dir; do
  name=$(basename "$dir")
  cd "./docker/$name"
  cp -r ../_base ./base
  docker build -t "${DOCKER_IMAGE_PREFIX}${name}" -q -f Dockerfile .
  cd "../.."
  echo "Built ${DOCKER_IMAGE_PREFIX}${name}"
  echo "$name" >> languages.txt
done

cd "$CURRENT_DIR"

docker image prune -f