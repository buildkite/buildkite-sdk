#!/bin/bash

set -eo pipefail

echo "INSTALLING GO SDK"

pushd sdk/go
go mod tidy
popd

echo "GO SDK INSTALLED"
exit 0
