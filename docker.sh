#!/bin/env bash

# Exist if any command fails
set -e

DOCKER_IMAGE="moabdelazem/automate-jenkins:latest"

# Check if the Docker image is already built
docker build -t $DOCKER_IMAGE .

# Run the container
docker run -p 8000:8000 --name automate-jenkins -d $DOCKER_IMAGE
