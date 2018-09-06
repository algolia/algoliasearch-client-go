#!/usr/bin/env bash
set -e

function is_version_valid() {
  echo "$1" | grep -q -E '^[0-9]+\.[0-9]+\.[0-9]+$'
  if [[ "$?" == "0" ]]; then return 0; else return 1; fi
}

function is_working_dir_clean() {
  if [ -z "$(git status --porcelain)" ]; then return 0; else return 1; fi
}

version=$1

if ! is_version_valid "$version"; then
  echo "Version '$version' is not valid (expecting X.Y.Z)"
  exit 1
fi

if ! is_working_dir_clean; then
  echo "Current directory is not clean, release aborted"
  exit 1
fi

gsed -i -E "s/version = \".+\"$/version = \"$version\"/" algoliasearch/transport.go
rake alg:changelog["$version"]

git --no-pager diff
printf 'Please confirm those final changes before the automatic release [y/n]: '
read yes_or_no
if [[ "$yes_or_no" != "y" ]]; then
  echo 'Aborting release'
  git reset --hard HEAD > /dev/null 2>&1
  exit 1
fi

git add ChangeLog.md algoliasearch/transport.go
git commit -m "chore: Release version $version [skip ci]"
git tag "$version"
git push --tags
