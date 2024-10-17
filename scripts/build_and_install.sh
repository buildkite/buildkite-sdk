#!/bin/bash

set -eo pipefail

echo "BUILDING SDKS"
rm -rf sdk
go run .

echo "INSTALL SDKS"
pushd sdk/go
go mod init github.com/buildkite/pipeline-sdk/sdk/go
go mod tidy
popd

pushd sdk/typescript
npm install
npm link
popd

echo "SDKS BUILT AND INSTALLED"