#!/usr/bin/env bash

set -e

target=${1?-Please provide target folder as first argument}

mkdir -p "${target}"
cp -r template/ "${target}"
sed --in-place "s,template,$target,g" "${target}/main.go"