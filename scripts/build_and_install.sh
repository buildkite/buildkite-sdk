#!/bin/bash

set -eo pipefail

go() {
    pushd sdk/go
    go mod init github.com/buildkite/pipeline-sdk/sdk/go
    go mod tidy
    popd
}

typescript() {
    pushd sdk/typescript
    npm install
    npm link
    popd
}

echo "BUILDING SDKS"
rm -rf sdk
go run .

if [ $1 = "go" ]; then
    go()
else if [ $1 = "typescript" ]; then
    typescript()
else
    go()
    typescript()
fi

echo "SDKS BUILT AND INSTALLED"