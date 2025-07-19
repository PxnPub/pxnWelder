#!/usr/bin/bash

\go get -u  || exit 1
\go build --buildmode=plugin -o weld-php.so .  || exit 1
