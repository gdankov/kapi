#!/bin/bash

readonly BASEDIR="$(cd "$(dirname "$0")" && pwd)"
readonly ROOTDIR="$BASEDIR/../"

pushd $ROOTDIR
kubectl apply -f ./crd/
pwd

popd
