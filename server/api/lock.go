//go:build !windows
// +build !windows

package api

import (
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func Lock() {
	var err error
	lock, err = os.Create(lockFile)
	if err != nil {
		logrus.Errorln("IOM:创建Lock失败", err)
	}
	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		Exit(1, "IOMServer重复执行，当前进程ID为:", os.Getpid(), ",退出.")
	}
}
func UnLock() {
	syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	lock.Close()
	os.Remove(lockFile)
}
