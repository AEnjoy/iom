package auth

import "time"

type loginInfo struct {
	Time     time.Time
	TryTimes int
}

var loginInfoMap = make(map[string]loginInfo) //ip:struct

func timeToClen() {
	for {
		for s, v := range loginInfoMap {
			nowT := time.Now().Unix()
			dataT := v.Time.Unix()
			if nowT-dataT > 600 {
				delete(loginInfoMap, s)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
