package devices

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetPackages
// http://127.0.0.1:8088/api/devices/:id/getpackages?token=xxx
// or
// http://127.0.0.1:8088/api/devices/:ClientID/getpackages?token=xxx
func GetPackages(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	id := context.Param("id") //int
	var clientID string
	idI, e := strconv.Atoi(id)
	if e != nil {
		//clinentToken
		clientID = id
		config.AddOperLogs("WebAPI访问：GetPackages", "GET", context.Request.URL.Path)
	} else {
		config.AddOperLogs("WebAPI访问：GetPackages", "POST", context.Request.URL.Path)
		clientID = config.DeviceIDGetDeviceToken(idI)
		if clientID == "" {
			context.String(http.StatusBadRequest, "Device is invalid")
			return
		}
	}
	context.JSON(http.StatusOK, global.PackagesInfo[clientID])
	logrus.Info("WebAPI:  Request ", context.Request.URL.Path, " Get Packages successfully.")
}
