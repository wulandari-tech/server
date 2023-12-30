#!/bin/bash -u
# This script compiles project for Linux amd64.
# It produces static C-libraries linkage.

wd=$(realpath -s "$(dirname "$0")/..")

cp -ruv "$wd/confdata" "$GOPATH/bin/config"

buildvers=$(git describe --tags)
buildtime=$(go run "$(dirname "$0")/timenow.go") # $(date -u +'%FT%TZ')

go env -w GOOS=linux GOARCH=amd64 CGO_ENABLED=1
go build -o $GOPATH/bin/slot_linux_x64 -v -ldflags="-linkmode external -extldflags -static -X 'github.com/slotopol/server/config.BuildVers=$buildvers' -X 'github.com/slotopol/server/config.BuildTime=$buildtime'" $wd
