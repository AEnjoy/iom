package global

import "time"
import "github.com/gin-gonic/gin"

var Router *gin.Engine

//var Read *bufio.Reader

var RPCPort = "" //default 10000
var WebPort = "" //default 8088

type TokenTimeToLife struct {
	Valid bool
	Time  time.Time
}

// 加密的Tokens
var TrustedTokens map[string]TokenTimeToLife

func TimeToCheckTokenIsValid() {
	for {
		for s, v := range TrustedTokens {
			nowT := time.Now().Unix()
			dataT := v.Time.Unix()
			if nowT-dataT > 60*60*24 { //一天
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

var DevicesInfos map[string]DevicesInfo //ClientID(Token),DevicesInfo
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
}

var PackagesInfo map[string]Packages
