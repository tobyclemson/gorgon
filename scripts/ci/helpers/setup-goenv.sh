#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"

eval "$(goenv init -)"

goenv install

GOROOT="$( goenv prefix )"
GOPATH="$HOME/go/$(cat .go-version)"

export GOROOT
export GOPATH

export PATH="$GOROOT/bin:$PATH"
export PATH="$GOPATH/bin:$PATH"
