#!/bin/bash

set -eo pipefail

echo "--- :pipeline: generating pipeline"
cd .buildkite
go run main.go

echo "--- :pipeline_upload: uploading pipeline.json"
buildkite-agent pipeline upload pipeline.json