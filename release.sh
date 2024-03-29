#!/usr/bin/env bash
set -e

version=$1

function is_version_valid() {
  echo "$1" | grep -q -E '^[0-9]+\.[0-9]+\.[0-9]+$'
  if [[ "$?" == "0" ]]; then return 0; else return 1; fi
}

function is_working_dir_clean() {
  if [ -z "$(git status --porcelain | grep -v -e ChangeLog.md)" ]; then return 0; else return 1; fi
}

function is_changelog_modified() {
  if [ -z "$(git diff ChangeLog.md | grep ${version})" ]; then
    echo "Missing ChangeLog update" >&2
    return 1
  else
    return 0
  fi
}

if ! is_version_valid "$version"; then
  echo "Version '$version' is not valid (expecting X.Y.Z)"
  exit 1
fi

if ! is_working_dir_clean; then
  echo "Current directory is not clean, release aborted"
  exit 1
fi
is_changelog_modified

gsed -i -E "s/version = \".+\"$/version = \"$version\"/" algolia/transport/transport.go


git --no-pager diff ChangeLog.md algolia/transport/transport.go
printf 'Please confirm those final changes before the automatic release [y/n]: '
read yes_or_no
if [[ "$yes_or_no" != "y" ]]; then
  echo 'Aborting release'
  git reset --hard HEAD > /dev/null 2>&1
  exit 1
fi

git add ChangeLog.md algolia/transport/transport.go
git commit -m "chore: Release version $version [skip ci]"
git push
git tag "v${version}"
git push origin "v${version}"
