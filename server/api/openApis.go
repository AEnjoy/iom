package api

import (
	"IOM/server/config"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var lock *os.File
var lockFile = "./IOMServerLock.pid"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

// 生成6位随机验证码（数字）
func Captcha1() (int, error) {
	return strconv.Atoi(fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)))
}

// 生成16位随机ID（字母）
func Capthca2() string {
	n := 16
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}