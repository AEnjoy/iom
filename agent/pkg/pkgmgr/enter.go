package pkgmgr

import (
	"runtime"
	"time"
	//pb "github.com/aenjoy/iom/agent/proto"
)

type rpmPackage struct {
	name        string
	version     string
	arch        string
	description string //null
	repository  string
}
type debPackage struct {
	name        string
	version     string
	arch        string
	description string //null
}
type apkPackage struct {
	name        string
	version     string
	arch        string
	description string //null
}

func Report() {
	if runtime.GOOS == "windows" {
		return
	}
	for {
		time.Sleep(time.Minute * 10)
		TimeToGetPackage()
	}
}

func GetPackages() {
	//var p pb.Package

}
