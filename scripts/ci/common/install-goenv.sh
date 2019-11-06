#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

git clone https://github.com/syndbg/goenv.git ~/.goenv
