#!/bin/bash

set -eo pipefail

FILES="$(git diff --name-only)"

if [ -z "${FILES}" ]; then
    exit 0
fi

echo "working tree is dirty"
git diff
exit 1
