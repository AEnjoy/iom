package devices

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type deviceInfos struct {
	total    int
	todayAdd int
}
type alarmInfos struct {
	total    int
	todayAdd int
}
type deviceLinkStatusInfos struct {
	linkStatus  string //online/offline
	deviceTotal int
}
type outData struct {
	deviceInfo           deviceInfos
	alarmInfo            alarmInfos
	deviceLinkStatusInfo []deviceLinkStatusInfos
	deviceCountType      []string
}
type alarmPanel struct {
	data  time.Time
	count int
}

// getPanels 获取面板信息
// http://127.0.0.1:8088/api/devices/get-panels
func getPanels(context *gin.Context) {
	logrus.Info("WebAPI: Request getPanels", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	var out outData
	out.deviceInfo.total = len(global.DevicesInfos)
	out.deviceInfo.todayAdd = config.GetTodayAddedDevicesMount()
	// Todo 告警数据未做
	out.alarmInfo.total = 0
	out.alarmInfo.todayAdd = 0
	out.deviceLinkStatusInfo = make([]deviceLinkStatusInfos, 0)
	var online, offline deviceLinkStatusInfos
	online.linkStatus = "online"
	online.deviceTotal = 0
	offline.linkStatus = "offline"
	offline.deviceTotal = 0
	for _, v := range global.DevicesInfos {
		if v.Online {
			online.deviceTotal++
		} else {
			offline.deviceTotal++
		}
	}
	out.deviceLinkStatusInfo = append(out.deviceLinkStatusInfo, online, offline)
	context.JSON(http.StatusOK, out)
	logrus.Info("WebAPI: Response getPanels", context.Request.URL.Path+" done.")
}

// getAlarmPanels 获取告警面板信息
// http://127.0.0.1:8088/api/devices/get-alarm-panels
func getAlarmPanels(context *gin.Context) {
	logrus.Info("WebAPI: Request getAlarmPanels", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	//var out outData
	type outData2 struct {
		data  time.Time
		count int
	}
	var out outData2
	out.count = 0
	for _, v := range global.AlarmInfos {
		if time.Now().Unix()-v.Time.Unix() < 60*60*24 {
			out.count++
		}
	}
	context.JSON(http.StatusOK, out)
	logrus.Info("WebAPI: Response getAlarmPanels", context.Request.URL.Path+" done.")
}
