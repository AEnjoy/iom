package main

import (
	"IOM/server/console"
	"IOM/server/global"
)

var (
	GoVersion string
	GitHash   string
	Version   string
	BuildTime string
)

func main() {
	global.Version = Version
	global.GitHash = GitHash
	global.GoVersion = GoVersion
	global.BuildTime = BuildTime
	console.Main()
}
