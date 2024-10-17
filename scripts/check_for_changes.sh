#!/bin/bash

set -eo pipefail

git ls-files --others --error-unmatch . >/dev/null 2>&1; ec=$?
if test "$ec" = 0; then
    echo "there are changes"
    exit 1
elif test "$ec" = 1; then
    exit 0
else
    echo "error from ls-files"
    exit 1
fi
