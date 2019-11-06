#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

set +e
openssl version
openssl aes-256-cbc \
    -d \
    -md sha1 \
    -in ./.circleci/gpg.private.enc \
    -k "${ENCRYPTION_PASSPHRASE}" | gpg --import -
set -e

git crypt unlock
