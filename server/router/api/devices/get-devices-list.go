package devices

import (
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// getDevices
// http://127.0.0.1:8088/api/devices/getdevices?token=xxx
func getDevices(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	type devicesList struct {
		Token string
		Type  int //0Windows 1Linux
	}
	var dt []devicesList
	for _, info := range global.DevicesInfos {
		if info.Online == true {
			var dtl devicesList
			if global.DevicesInfos[dtl.Token].HostInfo.Platform == "windows" {
				dtl.Type = 0
			} else if global.DevicesInfos[dtl.Token].HostInfo.Platform == "linux" {
				dtl.Type = 1
			}
			dtl.Token = info.ClientID
			dt = append(dt, dtl)
		}
	}
	if len(dt) > 0 {
		context.JSON(http.StatusOK, dt)
		logrus.Info("WebAPI:  Request ", context.Request.URL.Path, " Get Online Devices successfully.")
	} else {
		context.String(http.StatusNoContent, "No Online devices")
		logrus.Info("WebAPI:  Request ", context.Request.URL.Path, " Get Online Devices successfully. But no online devices")
	}
}
