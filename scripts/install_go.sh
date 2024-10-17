#!/bin/bash

set -eo pipefail

echo "INSTALLING GO SDK"

pushd sdk/go
ls
go mod init github.com/buildkite/pipeline-sdk/sdk/go
go mod tidy
popd

echo "GO SDK INSTALLED"
exit 0
