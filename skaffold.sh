#!/usr/bin/env bash

set -e

language="${1?-Please provide target languate as first argument}"
path="${2?-Please provide target path as first argument}"
folder="$(basename "$path")"

case $language in
  go)
    mkdir -p "${path}"
    cp 'template/main.go' "${path}/main.go"
    cp 'template/main_test.go' "${path}/main_test.go"
    sed --in-place "s,template,$path,g" "${path}/main.go"
    git add -A "${path}"
    ;;
  python)
    mkdir -p "${path}"
    cp 'template/dx.py' "${path}/d${folder}.py"
    cp 'template/dx_test.py' "${path}/d${folder}_test.py"
    ;;
  *)
    echo "Unsupported language: ${language}"
    exit 1
    ;;
esac

cp 'template/example.txt' "${path}"/example.txt
cp 'template/puzzle.txt' "${path}"/puzzle.txt
git add -A "${path}"