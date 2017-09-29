#!/usr/bin/env bash

GOPATH="$(pwd)/../../../../../"
GOOS=linux GOARCH=amd64 go build -v -o ../build/hlcup_linux_amd64 git.nulana.com/bobrnor/hlcup/hlcup
