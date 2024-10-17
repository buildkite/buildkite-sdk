#!/bin/bash

set -eo pipefail

FILES="$(git diff --name-only)"

if [ -z "${FILES}" ]; then
    exit 0
fi

echo "working tree is dirty"
DIFF="$(git diff)"
echo "${DIFF}"

exit 1
