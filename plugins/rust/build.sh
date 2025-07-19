#!/usr/bin/bash

\go get -u  || exit 1
\go build --buildmode=plugin -o weld-rust.so .  || exit 1
