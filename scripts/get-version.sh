#! /usr/bin/env bash
#
# current commit
commit=$(git rev-parse HEAD)
commit_short=$(git rev-parse --short HEAD)

last_tag=$(git describe --tags --abbrev=0)
last_tag_commit=$(git rev-list -n 1 $last_tag)

repo_root=$(git rev-parse --show-toplevel)
version_file=$repo_root/version.txt

is_formal=$1

if [ "$is_formal" == "true" ]; then
  printf $last_tag > $version_file
  exit 0
fi

if [ "$commit" == "$last_tag_commit" ]; then
  printf $last_tag > $version_file
else
  printf $last_tag-$commit_short > $version_file
fi

