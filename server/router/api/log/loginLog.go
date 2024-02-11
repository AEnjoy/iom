package log

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// ListLoginInfo
// http://127.0.0.1:8088/api/log/logLogin/list
func ListLoginInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request listLoginInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	var pageNum, pageSize int
	var loginLocation, username string
	pageNum, _ = strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ = strconv.Atoi(context.DefaultQuery("pageSize", "10"))
	loginLocation = context.DefaultQuery("loginLocation", "")
	username = context.DefaultQuery("username", "")
	logLogs := config.GetLogLogin(pageNum, pageSize, loginLocation, username)
	context.JSON(http.StatusOK, logLogs)
	logrus.Info("WebAPI: Request listLoginInfo", context.Request.URL.Path, "Done. Data:", logLogs)
}

// DelLoginInfo
// http://127.0.0.1:8088/api/log/logLogin/remove/:infoID
func DelLoginInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request delLoginInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	infoID, err := strconv.Atoi(context.Param("infoID"))
	if err != nil {
		context.String(http.StatusBadRequest, "infoID is invalid")
	}
	err = config.DelLoginInfo(infoID)
	if err != nil {
		context.String(http.StatusInternalServerError, "del login info failed.Error Info:", err)
	}
	context.String(http.StatusOK, "del login info success")
	logrus.Info("WebAPI: Request delLoginInfo", context.Request.URL.Path, "Done.")
}

// CleanAllLoginLog
// http://127.0.0.1:8088/api/log/logLogin/clean-all
func CleanAllLoginLog(context *gin.Context) {
	logrus.Info("WebAPI: Request CleanAllLoginLog", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	config.CleanAllLoginInfo()
	context.String(http.StatusOK, "clean all login info success")
}
