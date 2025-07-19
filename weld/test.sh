#!/usr/bin/bash
clear


pushd ../plugins/clean/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/deploy/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/gen/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/git/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/golang/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/gradle/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/make/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/maven/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/php/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/rpm/  || exit 1
	sh build.sh
popd >/dev/null

pushd ../plugins/rust/  || exit 1
	sh build.sh
popd >/dev/null


go mod tidy  || exit 1
go run . $@  || exit 1
