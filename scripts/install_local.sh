#!/bin/sh

set -e

NAME=showline

go build main.go
cp main ~/go/bin/$NAME