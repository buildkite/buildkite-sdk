#!/bin/bash

set -eo pipefail

echo "--- :pipeline: generating pipeline"
cd .buildkite
go run main.go