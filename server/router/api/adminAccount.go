package api

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func passwordChange(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	admin := context.PostForm("username")
	newPw := context.PostForm("password")
	err := config.AdminAccountEdit(admin, newPw)
	if err != nil {
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " ", err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.String(http.StatusOK, "ok")
}
