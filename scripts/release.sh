#! /usr/bin/env bash
#
# Create a new release.
# 1. Update the version in the version.txt file.
# 2. Update the version in the go version file under cmd/codetalks/version.go.
# 3. git commit the changes.
# 4. Create a new tag.
# 5. Push the changes to the remote.
# 6. Push the tag to the remote.
#

set -e

last_tag=$(git describe --tags --abbrev=0)
repo_root=$(git rev-parse --show-toplevel)
version_file=$repo_root/version.txt
go_version_file=$repo_root/cmd/codetalks/version.go

new_tag=$1

# A new tag must be provided.
if [ -z "$new_tag" ]; then
  echo "Error: A new tag must be provided. Current tag is: $last_tag"
  echo "Usage: $0 <new_tag>"
  exit 1
fi

# Check if tag already existed.
if $new_tag == $last_tag; then
  echo "Error: Tag $new_tag already exists."
  exit 0
fi

# 1. Update the version in the version.txt file.
printf $new_tag > $version_file

# 2. Update the version in the go version file.
version=$new_tag
echo "Updating new version: $version in go version file: $go_version_file"
echo "package main" > $go_version_file
echo "" >> $go_version_file
echo "const Version = \"$version\"" >> $go_version_file

# 3. git commit the changes.
git add .
git commit -m "New release: $new_tag"
git push

# 4. Create a new tag.
git tag $new_tag
git push origin $new_tag
