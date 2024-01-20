package global

import "time"

//var Read *bufio.Reader

var RPCPort = "" //default 10000
var WebPort = "" //default 8088

// 加密的Tokens
var TrustedTokens map[string]bool

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
	TrustedTokens = make(map[string]bool)
	PackagesInfo = make(map[string]Packages)
}

var PackagesInfo map[string]Packages
