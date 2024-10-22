#!/bin/bash

set -eo pipefail

process_go() {
    pushd sdk/go
    go mod init github.com/buildkite/pipeline-sdk/sdk/go
    go mod tidy
    popd
}

process_typescript() {
    pushd sdk/typescript
    npm install
    npm link
    popd
}

echo "BUILDING SDKS"
rm -rf sdk
go run .

case $1 in
    go)
        process_go
        ;;
    typescript)
        process_typescript
        ;;
    *)
        process_go
        process_typescript
        ;;
esac

echo "SDKS BUILT AND INSTALLED"
