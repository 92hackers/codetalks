#! /usr/bin/env bash

set -e

repo_root=$(git rev-parse --show-toplevel)
version_file=$repo_root/version.txt
go_version_file=$repo_root/cmd/codetalks/version.go

version=$(cat $version_file)

echo "Updating new version: $version in go version file: $go_version_file"

echo "package main" > $go_version_file
echo "" >> $go_version_file
echo "const Version=\"$version\"" >> $go_version_file

