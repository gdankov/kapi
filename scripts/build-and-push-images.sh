#!/bin/bash

readonly BASEDIR="$(cd "$(dirname "$0")" && pwd)"
readonly ROOTDIR="$BASEDIR/../"

pushd $ROOTDIR
docker build . -f cmd/bridge/Dockerfile -t "gdankov/kapi-bridge"
docker push "gdankov/kapi-bridge"
pwd

popd
