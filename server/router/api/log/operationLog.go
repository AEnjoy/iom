package log

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// CleanAllOpernfo
// http://127.0.0.1:8088/api/log/logOper/all
func CleanAllOpernfo(context *gin.Context) {
	logrus.Info("WebAPI: Request CleanAllOpernfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	config.CleanAllOperLogs()
	context.String(http.StatusOK, "success")
	logrus.Info("WebAPI: Request ", context.Request.URL.Path, " success.")
}

// DelOperInfo
// http://127.0.0.1:8088/api/log/logOper/remove/:id
func DelOperInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request DelOperInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	id, _ := strconv.Atoi(context.Param("id"))
	config.DelOperLogs(id)
	context.String(http.StatusOK, "success")
	logrus.Info("WebAPI: Request ", context.Request.URL.Path, " success.")
}

// ListOperInfo
// http://127.0.0.1:8088/api/log/logOper/list?pageNo=xxx&pageSize=xxx&title=xxx&operName=xxx&businessType=xxx
func ListOperInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request ListOperInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	pageNo, _ := strconv.Atoi(context.DefaultQuery("pageNo", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))
	title := context.DefaultQuery("title", "")
	operName := context.DefaultQuery("operName", "")
	businessType := context.DefaultQuery("businessType", "")
	context.JSON(http.StatusOK, config.GetOperInfo(pageNo, pageSize, title, operName, businessType))
	logrus.Info("WebAPI: Request ", context.Request.URL.Path, " success.")
}
