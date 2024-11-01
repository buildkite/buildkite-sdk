#!/bin/bash

set -eo pipefail

docker build docker build --platform linux/amd64 -t zchase399/buildkite-pipeline-sdk-build .

docker tag buildkite/pipeline-sdk-build:latest zchase399/buildkite-pipeline-sdk-build:0.0.2

docker push zchase399/buildkite-pipeline-sdk-build:0.0.2
