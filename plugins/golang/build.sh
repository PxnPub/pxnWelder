#!/usr/bin/bash

\go get -u  || exit 1
\go build --buildmode=plugin -o weld-golang.so .  || exit 1
