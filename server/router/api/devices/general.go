package devices

import (
	"IOM/server/api"
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// deviceInfo  http://127.0.0.1:8088/api/devices/:id/info?token=xxx
//
//	or	 http://127.0.0.1:8088/api/devices/:ClientID/info?token=xxx
//	json
func deviceInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	id := context.Param("id") //int
	var clinentID string
	idI, e := strconv.Atoi(id)
	if e != nil {
		//clinentToken
		clinentID = id
	} else {
		clinentID = config.DeviceIDGetDeviceToken(idI)
		if clinentID == "" {
			context.String(http.StatusBadRequest, "DeviceID is invalid")
			return
		}
	}
	//v, e := json.Marshal(global.DevicesInfos[clinentID])
	context.JSON(http.StatusOK, global.DevicesInfos[clinentID])
}

// add http://127.0.0.1:8088/api/devices/add?token=xxx&
//
//	deviceId=xxx&devicesWeight=xxx&devicesToken=xxx&deviceName=xxx&deviceFlag=xxx&groupId=xxx
func add(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	deviceId, _ := strconv.Atoi(context.PostForm("deviceId"))
	devicesWeight, _ := strconv.Atoi(context.PostForm("devicesWeight"))
	devicesToken := context.PostForm("devicesToken")
	deviceName := context.PostForm("deviceName")
	deviceFlag, _ := strconv.Atoi(context.PostForm("deviceFlag"))
	groupId, _ := strconv.Atoi(context.PostForm("groupId"))
	if deviceId == 0 {
		var t int
		for {
			t, _ = api.Captcha1()
			if !config.IsDeviceIDValid(t) {
				break
			}
		}
	}
	logrus.Info("WebAPi: add deviceId:", deviceId, " ,devicesWeight:", devicesWeight, " ,devicesToken:", devicesToken, " ,deviceName", deviceName, " ,deviceFlag:", deviceFlag, " ,groupId:", groupId)
	err := config.DeviceAdd(deviceId, devicesWeight, devicesToken, deviceName, deviceFlag, groupId)
	if err != nil {
		context.String(http.StatusInternalServerError, "add device failed. Info:", err.Error())
		return
	}
	var tokenApi config.TokenApi
	tokenApi.SetClientToken(devicesToken)
	t, _ := tokenApi.GetTrustToken()
	global.TrustedTokens[t] = global.TokenTimeToLife{Valid: true, Time: time.Now()}

	logrus.Info("WebAPI: add trusted token:", t)
	context.String(http.StatusOK, "ok")
}

// deleteDevice
// http://127.0.0.1:8088/api/devices/delete?token=xxx&id=xxx&groupid=xxx
func deleteDevice(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	deviceId, _ := strconv.Atoi(context.DefaultQuery("id", ""))
	groupID, _ := strconv.Atoi(context.DefaultQuery("groupid", ""))
	logrus.Info("WebAPi: delete deviceId:", deviceId, " ,groupId:", groupID)
	err := config.DeviceDelete(deviceId, groupID)
	if err != nil {
		context.String(http.StatusInternalServerError, "delete device failed. Info:", err.Error())
		return
	}
	context.String(http.StatusOK, "ok")
}
