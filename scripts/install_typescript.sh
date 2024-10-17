#!/bin/bash

set -eo pipefail

echo "INSTALLING GO SDK"

pushd sdk/typescript
npm install
npm link
popd

echo "GO SDK INSTALLED"
exit 0