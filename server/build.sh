#!/usr/bin/env bash
# windows/amd64,386,amr64,arm32
# linux/amd64,386,486,arm32
#
#
#
#
#
Version=$(cat version)
COMMIT=$(git rev-parse --short HEAD)
GoVersion=$(echo $(go version) | awk '{print $3}')
Time=$(date "+%Y-%m-%d:::%H:%M:%S")
export GOOS=linux
export GOARCH=386
export CGO_LDFLAGS="-static"

go build  -v -ldflags "-s -w -X main.Version=$Version -X main.GitHash=$COMMIT -X main.GoVersion=$GoVersion -X main.BuildTime=$Time"  ./main.go