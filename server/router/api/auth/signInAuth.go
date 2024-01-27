package auth

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Get
func signInAuth(context *gin.Context) {
	if global.IsCookieAuthValid(context) { //防止反复登录授权
		context.String(http.StatusOK, "ok")
		//context.Redirect(302, "/dashboard")
		return
	}
	un := context.DefaultQuery("username", "")
	pw := context.DefaultQuery("password", "")
	var tokenApi config.TokenApi
	tokenApi.SetPassword(pw)
	tokenApi.SetUserName(un)
	t, e := tokenApi.GetTrustToken()
	if e != nil {
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, e.Error())
		return
	}
	//bug：t为空时全取true
	_, ok := global.TrustedTokens[t]
	//println(t, v, ok)
	if ok {
		context.SetCookie("IOMAuth", t, 3600*24, "", context.Request.Host, false, false)
		//context.Redirect(302, "/dashboard")
		context.String(http.StatusOK, "Verified")
		return
	}
	context.String(http.StatusUnauthorized, "Verify failed. Username or password is not correct?")
	logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw)
}

// Post
// curl http://localhost:8088/api/auth/signin  -X POST -d 'username=admin&password=passwords'
func signInAuthP(context *gin.Context) {
	if global.IsCookieAuthValid(context) { //防止反复登录授权
		//context.String(http.StatusOK, "ok")
		context.String(http.StatusOK, "Verified")
		return
	}
	//println(context.Request.Body)
	un := context.PostForm("username")
	pw := context.PostForm("password")
	var tokenApi config.TokenApi
	tokenApi.SetPassword(pw)
	tokenApi.SetUserName(un)
	t, e := tokenApi.GetTrustToken()
	if e != nil {
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, e.Error())
		return
	}
	//bug：t为空时全取true
	_, ok := global.TrustedTokens[t]
	//println(t, v, ok)
	if ok {
		context.SetCookie("IOMAuth", t, 3600*24, "", context.Request.Host, false, false)
		//context.Redirect(302, "/dashboard")
		context.String(http.StatusOK, "Verified")
		return
	}
	context.String(http.StatusUnauthorized, "Verify failed. Username or password is not correct?")
	logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw)
}
