#!/usr/bin/bash

\go get -u  || exit 1
\go build --buildmode=plugin -o weld-make.so .  || exit 1
