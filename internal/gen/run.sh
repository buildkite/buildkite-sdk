#!/bin/bash

set -o errexit -o pipefail

rm -rf definitions
mkdir definitions
go run .
~/go/bin/goimports -w ./definitions
gofmt ./definitions