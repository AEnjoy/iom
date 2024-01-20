package global

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IsValidToken(token string) bool {
	if token == "" {
		return false
	} //空token的map会返回true
	_, ok1 := DevicesInfos[token]
	_, ok2 := TrustedTokens[token]
	//Debug
	return ok1 || ok2
}

func IsCookieAuthValid(c *gin.Context) bool {
	cookie, err := c.Cookie("IOMAuth") //加密的
	if err != nil {
		logrus.Errorln("Cookie err: " + err.Error())
	}
	logrus.Debugln("API IsCookieAuthValid: get cookie: " + cookie)
	return IsValidToken(cookie)
}
