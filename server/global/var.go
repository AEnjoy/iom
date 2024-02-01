package global

import (
	"github.com/sirupsen/logrus"
	"time"
)
import "github.com/gin-gonic/gin"

var Router *gin.Engine
var DebugFlag = true

//var Read *bufio.Reader

type TokenTimeToLife struct {
	Valid bool
	Time  time.Time
}

// 加密的Tokens
var TrustedTokens map[string]TokenTimeToLife
var AlarmInfos []AlarmInfo

func TimeToCheckTokenIsValid() {
	for {
		for s, v := range TrustedTokens {
			nowT := time.Now().Unix()
			dataT := v.Time.Unix()
			if nowT-dataT > 60*60*24 { //一天
				logrus.Info("Token is expired:", s)
				delete(TrustedTokens, s)
			}
		}
		time.Sleep(time.Second * 10)
	}
}
func SetDeviceOnline(token string) {
	v, ok := DevicesInfos[token]
	if v.Online == false && ok == true {
		dt := DevicesInfos[token]
		dt.Online = true
		DevicesInfos[token] = dt
	}
}

var DevicesInfos map[string]DevicesInfo //ClientID(Token),DevicesInfo 所有设备列表
func TimeToCheckOnline() {
	for {
		for s, info := range DevicesInfos {
			nowT := time.Now().Unix()
			dataT := info.DataTime.Unix()
			if nowT-dataT >= 15 { //second
				var dt DevicesInfo
				dt = info
				dt.Online = false
				DevicesInfos[s] = dt
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func MapInit() {
	DevicesInfos = make(map[string]DevicesInfo)
	TrustedTokens = make(map[string]TokenTimeToLife)
	PackagesInfo = make(map[string]Packages)
	AlarmInfos = make([]AlarmInfo, 0)
}

var PackagesInfo map[string]Packages
