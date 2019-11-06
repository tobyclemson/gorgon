#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/../../.." && pwd )"

cd "$PROJECT_DIR"

set +e
openssl version
openssl aes-256-cbc \
    -d \
    -md sha1 \
    -in ./.circleci/gpg.private.enc \
    -k "${ENCRYPTION_PASSPHRASE}" | gpg --import -
set -e

git crypt unlock

export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"

eval "$(goenv init -)"

goenv install

export GOROOT="$( goenv prefix )"
export GOPATH="$HOME/go/$(cat .go-version)"

export PATH="$GOROOT/bin:$PATH"
export PATH="$GOPATH/bin:$PATH"

gem install bundler
bundle install

rake
