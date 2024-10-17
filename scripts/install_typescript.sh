#!/bin/bash

set -eo pipefail

echo "INSTALLING TYPESCRIPT SDK"

pushd sdk/typescript
npm install
npm link
popd

echo "TYPESCRIPT SDK INSTALLED"
exit 0