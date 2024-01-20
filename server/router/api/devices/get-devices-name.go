package devices

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// getDevicesName
// ttp://127.0.0.1:8088/api/devices/:devices-token/getdevices-name?token=xxx
func getDevicesName(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	id := context.Param("id")
	name, err := config.DeviceTokenGetDeviceName(id)
	if err != nil {
		context.String(http.StatusNotAcceptable, "Device is invalid")
		return
	}
	context.String(http.StatusOK, name)
	logrus.Info("WebAPI:  Request ", context.Request.URL.Path, " Get DevicesName successfully.")
}
