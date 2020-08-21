#!/bin/bash

docker build . -t precommit -f preCommit.Dockerfile
docker run -it -v "$(pwd):/app" -v "/tmp/.cache:/root/.cache" precommit
