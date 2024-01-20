//go:build windows

package api

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Lock() {
	var err error
	_, err = os.Stat(lockFile)
	if err != nil && os.IsNotExist(err) {
		lock, err = os.Create(lockFile)
		if err != nil {
			logrus.Errorln("IOM:创建Lock失败", err)
		}
	} else {
		Exit(1, "IOMServer重复执行，当前进程ID为:", os.Getpid(), ",退出.")
	}
}
func UnLock() {
	lock.Close()
	os.Remove(lockFile)
}
