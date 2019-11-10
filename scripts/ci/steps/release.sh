#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/../../.." && pwd )"

cd "$PROJECT_DIR"

source "./scripts/ci/helpers/setup-goenv.sh"
source "./scripts/ci/helpers/unlock-git-crypt.sh"
source "./scripts/ci/helpers/install-gems.sh"

rake cli:release
git fetch
rake homebrew:formula:push
