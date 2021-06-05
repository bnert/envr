#!/usr/bin/env bash

mkdir -p build

TOP_LVL_DIR=$(git rev-parse --show-toplevel)

function bin {
  go build -o build/envr cmd/envr.go
}

pushd ${TOP_LVL_DIR} > /dev/null
$1
popd > /dev/null
