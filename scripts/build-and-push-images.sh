#!/bin/bash

readonly BASEDIR="$(cd "$(dirname "$0")" && pwd)"
readonly ROOTDIR="$BASEDIR/../"

pushd $ROOTDIR
docker build . -f cmd/bridge/Dockerfile -t "gdankov/kapi-bridge"
docker build . -f cmd/controller/lrp/Dockerfile -t "gdankov/lrp-controller"
docker build . -f cmd/controller/staging/Dockerfile -t "gdankov/staging-controller"

docker push "gdankov/kapi-bridge"
docker push "gdankov/lrp-controller"
docker push "gdankov/staging-controller"
pwd

popd
