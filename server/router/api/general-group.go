package api

/*
GroupAPIs
*/

import (
	"IOM/server/api"
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// groupAdd
// Post
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
	gid, e := strconv.Atoi(context.PostForm("id"))
	if gid == 0 {
		var t int
		for {
			t, _ = api.Captcha1()
			if !config.IsGroupIDValid(t) {
				break
			}
		}
		gid = t
	}
	if e != nil {
		context.String(http.StatusBadRequest, "id is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " id is invalid.")
		return
	}
	name := context.PostForm("name")
	err := config.GroupAdd(gid, name)
	if err != nil {
		context.String(http.StatusInternalServerError, "group add failed")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " group add failed.")
		return
	}
	context.String(http.StatusOK, "ok")
	logrus.Infoln("WebAPI: group add id:", gid, ", name:", name, " successful.")
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

//	     groupIDGetGroupName Get/Post
//			http://localhost:8088/api/group/id/is-valid?id=xxx
func isGroupIDValidGetMethod(context *gin.Context) {
	logrus.Info("WebAPI: Request isGroupIDValidGetMethod:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request Delete:", context.Request.URL.Path, " token is invalid.")
		return
	}
	gId, _ := strconv.Atoi(context.DefaultQuery("id", "-1"))
	context.String(http.StatusOK, strconv.FormatBool(config.IsGroupIDValid(gId)))
	logrus.Infoln("WebAPI: Request isGroupIDValidPostMethod:", context.Request.URL.Path, " done. Value is ", strconv.FormatBool(config.IsGroupIDValid(gId)))
}
func isGroupIDValidPostMethod(context *gin.Context) {
	logrus.Info("WebAPI: Request isGroupIDValidGetMethod:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request Delete:", context.Request.URL.Path, " token is invalid.")
		return
	}
	gId, _ := strconv.Atoi(context.PostForm("id"))
	context.JSON(http.StatusOK, gin.H{"Status": strconv.FormatBool(config.IsGroupIDValid(gId))})
	logrus.Infoln("WebAPI: Request isGroupIDValidPostMethod:", context.Request.URL.Path, " done. Value is ", strconv.FormatBool(config.IsGroupIDValid(gId)))
}

// getGroupList
// curl http://localhost:8088/api/group/get-list
func getGroupList(context *gin.Context) {
	logrus.Info("WebAPI: Request getGroupList:", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request get:", context.Request.URL.Path, " token is invalid.")
		return
	}
	type group struct {
		GroupID   int
		GroupName string
	}
	var groupList []group
	for _, info := range config.DevicesGroupEX {
		var t group
		t.GroupID = config.GroupNameGetGroupID(info.GroupName)
		t.GroupName = info.GroupName
		groupList = append(groupList, t)
	}
	context.JSON(http.StatusOK, groupList)
	logrus.Infoln("WebAPI: Request getGroupList:", context.Request.URL.Path, " done.")
}
