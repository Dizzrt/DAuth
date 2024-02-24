#!/bin/sh

set -e

$VERSION = git describe --tags --always
# $GIT_COMMIT = $( git rev-parse --short HEAD)
# $BUILD_TIME = $( date -R)

define LDFLAGS
"-X 'github.com/myname/myapp/config.Version=${VERSION}' \
-X 'github.com/myname/myapp/config.GitCommit=${GIT_COMMIT}' \
-X 'github.com/myname/myapp/config.BuildTime=${BUILD_TIME}'"
endef

build:
    go build -ldflags ${LDFLAGS} -o myapp_${VERSION} .
