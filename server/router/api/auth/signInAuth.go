package auth

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Get
func SignInAuth(context *gin.Context) {
	if global.IsCookieAuthValid(context) { //防止反复登录授权
		v, _ := context.Cookie("IOMAuth")
		context.String(http.StatusOK, v)
		//context.Redirect(302, "/dashboard")
		return
	}
	un := context.DefaultQuery("username", "")
	pw := context.DefaultQuery("password", "")
	v, ok := loginInfoMap[context.Request.RemoteAddr]
	if v.TryTimes > 4 {
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?There are too many login failure errors, please try again in 10 minutes.")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, "登录失败次数过多")
		return
	}
	var tokenApi config.TokenApi
	tokenApi.SetPassword(pw)
	tokenApi.SetUserName(un)
	t, e := tokenApi.GetTrustToken()
	if e != nil {
		if ok {
			v.TryTimes++
		} else {
			loginInfoMap[context.Request.RemoteAddr] = loginInfo{Time: time.Now(), TryTimes: 1}
		}
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, e.Error())
		return
	}
	//bug：t为空时全取true
	_, ok = global.TrustedTokens[t]
	//println(t, v, ok)
	if ok {
		context.SetCookie("IOMAuth", t, 3600*24, "", context.Request.Host, false, false)
		context.SetCookie("username", un, 3600*24, "", context.Request.Host, false, false)
		config.AddLoginLog(un, context.Request.RemoteAddr, context.Request.Host, context.Request.UserAgent(), "登录成功", time.Now())
		//context.Redirect(302, "/dashboard")
		context.String(http.StatusOK, t)
		return
	}
	context.String(http.StatusUnauthorized, "Verify failed. Username or password is not correct?")
	logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw)
}

// Post
// curl http://localhost:8088/api/auth/signin  -X POST -d 'username=admin&password=passwords'
func SignInAuthP(context *gin.Context) {
	if global.IsCookieAuthValid(context) { //防止反复登录授权
		//context.String(http.StatusOK, "ok")
		v, _ := context.Cookie("IOMAuth")
		context.String(http.StatusOK, v)
		return
	}
	//println(context.Request.Body)
	un := context.PostForm("username")
	pw := context.PostForm("password")
	v, ok := loginInfoMap[context.Request.RemoteAddr]
	if v.TryTimes > 4 {
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?There are too many login failure errors, please try again in 10 minutes.")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, "登录失败次数过多")
		return
	}
	var tokenApi config.TokenApi
	tokenApi.SetPassword(pw)
	tokenApi.SetUserName(un)
	t, e := tokenApi.GetTrustToken()
	if e != nil {
		if ok {
			v.TryTimes++
		} else {
			loginInfoMap[context.Request.RemoteAddr] = loginInfo{Time: time.Now(), TryTimes: 1}
		}
		context.String(http.StatusUnauthorized, "sign in failed. Username or password is not correct?")
		logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw, e.Error())
		return
	}
	//bug：t为空时全取true
	_, ok = global.TrustedTokens[t]
	//println(t, v, ok)
	if ok {
		context.SetCookie("IOMAuth", t, 3600*24, "", context.Request.Host, false, false)
		context.SetCookie("username", un, 3600*24, "", context.Request.Host, false, false)
		config.AddLoginLog(un, context.Request.RemoteAddr, context.Request.Host, context.Request.UserAgent(), "登录成功", time.Now())
		//context.Redirect(302, "/dashboard")
		context.String(http.StatusOK, t)
		return
	}
	context.String(http.StatusUnauthorized, "Verify failed. Username or password is not correct?")
	logrus.Error("失败的登录请求：From ", context.Request.RemoteAddr, ", username:", un, ", password:", pw)
}
