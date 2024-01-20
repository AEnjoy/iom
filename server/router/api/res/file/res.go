package file

import (
	serviceApi "IOM/server/api"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Main http://127.0.0.1:8088/api/res/file/whatever?token=xxxx
func Main() {
	global.Router.GET("/api/res/file/:file", func(context *gin.Context) {
		logrus.Info("WebAPI: Request ", context.Request.URL.Path)
		token := context.DefaultQuery("token", "")
		if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
			context.String(http.StatusUnauthorized, "token is invalid")
			logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
			return
		}
		filename := context.Param("file")
		t := serviceApi.FileReadAll("./res/scripts/" + filename)
		if t == "" {
			context.String(http.StatusNotFound, "file not found")
			return
		}
		context.String(http.StatusOK, t)
	})
}
