#!/usr/bin/bash
clear

pushd plugins/clean/ >/dev/null || exit 1
	sh build.sh
popd >/dev/null

#pushd plugins/git/ >/dev/null || exit 1
#	sh build.sh
#popd >/dev/null

#pushd plugins/golang/ >/dev/null || exit 1
#	sh build.sh
#popd >/dev/null

pushd plugins/rpm/ >/dev/null || exit 1
	sh build.sh
popd >/dev/null


go mod tidy  || exit 1
go run . $@  || exit 1
