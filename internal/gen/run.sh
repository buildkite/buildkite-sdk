#!/bin/bash

set -o errexit -o pipefail

go run .
~/go/bin/goimports -w ./definitions
gofmt ./definitions