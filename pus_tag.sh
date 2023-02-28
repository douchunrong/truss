#!/bin/zsh
set -x
version=$1
desc=$2
if [ "${desc}" == "" ]; then
  desc="create new version ${version}"
fi
git tag -a v${version} -m $desc
git push origin --tags
