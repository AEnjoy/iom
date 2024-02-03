#!/usr/bin/env bash

Version=$(cat version)
COMMIT=$(git rev-parse --short HEAD)
GoVersion=$(echo $(go version) | awk '{print $3}')
Time=$(date "+%Y-%m-%d:::%H:%M:%S")
export GOOS=windows
export GOARCH=amd64

go build -ldflags "-X main.Version=$Version -X main.GitHash=$COMMIT -X main.GoVersion=$GoVersion -X main.BuildTime=$Time"  ./main.go