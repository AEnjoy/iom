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
// Get
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

//	       groupNameGetGroupID Get
//		   	http://127.0.0.1:8088/api/group/search/name-get-id?token=xxx&name=xxx
func groupNameGetGroupIDGetMethod(context *gin.Context) {
	logrus.Info("WebAPI: Request GetGroupID:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request Delete:", context.Request.URL.Path, " token is invalid.")
		return
	}
	name := context.DefaultQuery("name", "")
	gId := config.GroupNameGetGroupID(name)
	if gId == 0 {
		context.String(http.StatusNotFound, "group not found")
		logrus.Infoln("WebAPI: Request ", context.Request.URL.Path, " done, but group not found.")
		return
	}
	context.String(http.StatusOK, strconv.Itoa(gId))
	logrus.Infoln("WebAPI: Request ", context.Request.URL.Path, " done, group id:", gId)
}

//	 groupNameGetGroupID Post
//		curl http://localhost:8088/api/group/search/name-get-id  -X POST -d 'name=xxx'
func groupNameGetGroupIDPostMethod(context *gin.Context) {
	logrus.Info("WebAPI: Request GetGroupID:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request Delete:", context.Request.URL.Path, " token is invalid.")
		return
	}
	name := context.PostForm("name")
	gId := config.GroupNameGetGroupID(name)
	if gId == 0 {
		context.JSON(http.StatusNotFound, gin.H{"ID": 0, "Status": "Not Found"})
		logrus.Infoln("WebAPI: Request ", context.Request.URL.Path, " done, but group not found.")
		return
	}
	context.JSON(http.StatusOK, gin.H{"ID": gId, "Status": "OK"})
	logrus.Infoln("WebAPI: Request ", context.Request.URL.Path, " done, group id:", gId)
}
