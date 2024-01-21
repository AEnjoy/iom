package api

import (
	"IOM/server/config"
	"bufio"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var lock *os.File
var lockFile = "./IOMServerLock.pid"

func Exit(code int, msg ...interface{}) {
	if code != 0 {
		logrus.Errorln("ServerExitWithMessage:", msg)
	} else {
		logrus.Info("ServerExitWithMessage:", msg)
	}
	logrus.Info("ServerExitWithCode:", code)
	UnLock()
	os.Exit(code)
}
func FileReadAll(filename string) string {
	var str = ""
	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			str = str + s
		}
	}
	return str
}
func ExitHandle(exitChan chan os.Signal) {
	for {
		select {
		case sig := <-exitChan:
			logrus.Info("收到来自系统的信号：", sig)
			config.DBExitSave()
			Exit(2, sig.String())
		}
	}

}
