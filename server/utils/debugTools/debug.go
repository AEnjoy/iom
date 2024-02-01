package debugTools

import (
	"IOM/server/global"
	"github.com/sirupsen/logrus"
)

func PrintLogs(a ...interface{}) {
	if global.DebugFlag {
		logrus.Debugln(a)
	} else {
		logrus.Infoln(a)
	}
}

func Println(a ...interface{}) {
	if global.DebugFlag {
		println("Debug:", a)
	} else {
		println(a)
	}
}
func PrintlnOnlyInDebugMode(a ...interface{}) bool {
	if global.DebugFlag {
		Println(a)
	}
	return global.DebugFlag
}
func PrintLogsOnlyInDebugMode(a ...interface{}) bool {
	if global.DebugFlag {
		PrintLogs(a)
	}
	return global.DebugFlag // 返回debug flag
}
