#!/bin/bash

set -eo pipefail

echo "BUILDING SDKS"
rm -rf sdk
go run .