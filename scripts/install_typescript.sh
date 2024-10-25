#!/bin/bash

set -eo pipefail

echo "INSTALLING TYPESCRIPT SDK"

pushd sdk/typescript
npx tsc --init --outDir ./dist
npm install
npm run build
npm link
popd

echo "TYPESCRIPT SDK INSTALLED"
exit 0