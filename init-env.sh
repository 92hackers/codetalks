#! /usr/bin/env bash
#
# Init local environment for development.
# 1. Install go dependencies.
# 2. Copy the pre-commit hook to the .git/hooks directory.
#

set -e

# Install go dependencies.
go mod tidy && go mod download

# Copy the pre-commit hook to the .git/hooks directory.
cp -f scripts/pre-commit .git/hooks/pre-commit
