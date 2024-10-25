#!/bin/bash

set -eo pipefail

test_go() {
    pushd sdk_tests/go
    go test .
    popd
}

test_typescript() {
    pushd sdk/typescript
    npm install
    npm run build
    npm link
    popd

    pushd sdk_tests/typescript
    npm install
    npm run test
    popd
}

echo "TESTING SDKS"

case $1 in
    go)
        test_go
        ;;
    typescript)
        test_typescript
        ;;
    *)
        test_go
        test_typescript
        ;;
esac

echo "SDKS TESTED"
