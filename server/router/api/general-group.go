package api

/*
GroupAPIs
*/

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// groupAdd
//
//	http://127.0.0.1:8088/api/group/add?token=xxx&id=xxx&name=xxx
func groupAdd(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}

	id, e := strconv.Atoi(context.PostForm("id"))
	if e != nil {
		context.String(http.StatusBadRequest, "id is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " id is invalid.")
		return
	}
	name := context.PostForm("name")
	err := config.GroupAdd(id, name)
	if err != nil {
		context.String(http.StatusInternalServerError, "group add failed")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " group add failed.")
		return
	}
	context.String(http.StatusOK, "ok")
	logrus.Infoln("WebAPI: group add id:", id, ", name:", name, " successful.")
}

//		groupDelete
//	 http://127.0.0.1:8088/api/group/delete?token=xxx&id=xxx
func groupDelete(context *gin.Context) {
	logrus.Info("WebAPI: Request Delete:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request Delete:", context.Request.URL.Path, " token is invalid.")
		return
	}
	id, e := strconv.Atoi(context.DefaultQuery("id", ""))
	if e != nil {
		context.String(http.StatusBadRequest, "id is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " id is invalid.")
		return
	}
	err := config.GroupDelete(id)
	if err != nil {
		context.String(http.StatusInternalServerError, "group delete failed")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " group delete failed.")
		return
	}
	context.String(http.StatusOK, "ok")
	logrus.Infoln("WebAPI: group delete id:", id, " successful.")
}
