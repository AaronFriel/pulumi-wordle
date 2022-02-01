#!/bin/bash
set -o nounset -o errexit -o pipefail

# ensure the index is up to date, in CI we were seeing some cases
# where git diff-files would show the entire tree was unmerged
# and this addresses that.
git update-index -q --refresh

if ! git diff-files --quiet; then
    >&2 echo "error: working tree is not clean, aborting!"
    git status
    git diff
    exit 1
fi
