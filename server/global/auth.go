package global

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IsValidToken(token string) bool {
	if token == "" {
		return false
	} //空token的map会返回true
	temp, ok2 := TrustedTokens[token]
	if DebugFlag { //当使用调试模式时，将允许直接使
		// 用设备token授权
		logrus.Debugln("Debug mode, allow access with device token: " + token)
		_, ok1 := DevicesInfos[token]
		return ok1 || (ok2 && temp.Valid)
	}
	return ok2 && temp.Valid
}

func IsCookieAuthValid(c *gin.Context) bool {
	cookie, err := c.Cookie("IOMAuth") //加密的
	if err != nil {
		logrus.Errorln("Cookie err: " + err.Error())
	}
	logrus.Debugln("API IsCookieAuthValid: get cookie: " + cookie)
	return IsValidToken(cookie)
}
