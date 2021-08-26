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

export GIT_SSH_COMMAND="ssh -i ~/.ssh/id_rsa_63092536911eea685e909417d5a5dacb"

bundle exec rake cli:release
bundle exec rake homebrew:formula:push
