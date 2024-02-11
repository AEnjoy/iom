package devices

import (
	"IOM/server/config"
	"IOM/server/global"
	"IOM/server/utils/debugTools"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type deviceInfos struct {
	Total    int
	TodayAdd int
}
type alarmInfos struct {
	Total    int
	TodayAdd int
}
type deviceLinkStatusInfos struct {
	LinkStatus  string //online/offline
	DeviceTotal int
}
type outData struct {
	DeviceInfo           deviceInfos
	AlarmInfo            alarmInfos
	DeviceLinkStatusInfo []deviceLinkStatusInfos
	DeviceCountType      []string
}
type alarmPanel struct {
	Data  time.Time
	Count int
}

// GetPanels 获取面板信息
// http://127.0.0.1:8088/api/devices/get-panels
func GetPanels(context *gin.Context) {
	logrus.Info("WebAPI: Request getPanels", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	var out outData
	out.DeviceCountType = make([]string, 0)
	out.DeviceInfo.Total = len(global.DevicesInfos)
	out.DeviceInfo.TodayAdd = config.GetTodayAddedDevicesMount()
	// Todo 告警数据未做
	out.AlarmInfo.Total = 0
	out.AlarmInfo.TodayAdd = 0
	out.DeviceLinkStatusInfo = make([]deviceLinkStatusInfos, 0)
	var online, offline deviceLinkStatusInfos
	online.LinkStatus = "online"
	online.DeviceTotal = 0
	offline.LinkStatus = "offline"
	offline.DeviceTotal = 0
	for _, v := range global.DevicesInfos {
		if v.Online {
			online.DeviceTotal++
		} else {
			offline.DeviceTotal++
		}
	}
	out.DeviceLinkStatusInfo = append(out.DeviceLinkStatusInfo, online, offline)
	debugTools.PrintLogs("Debug: getPanels 输出out")
	debugTools.PrintLogs(out.DeviceInfo.Total)
	debugTools.PrintLogs(out.DeviceCountType)
	debugTools.PrintLogs(out.DeviceLinkStatusInfo)
	debugTools.PrintLogs(out.AlarmInfo.Total)
	debugTools.PrintLogs(out.AlarmInfo.TodayAdd)
	context.JSON(http.StatusOK, out)
	logrus.Info("WebAPI: Response getPanels", context.Request.URL.Path+" done.")
}

// GetAlarmPanels 获取告警面板信息
// http://127.0.0.1:8088/api/devices/get-alarm-panels
func GetAlarmPanels(context *gin.Context) {
	logrus.Info("WebAPI: Request getAlarmPanels", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	//var out outData
	type outData2 struct {
		Data  time.Time
		Count int
	}
	var out outData2
	out.Count = 0
	for _, v := range global.AlarmInfos {
		if time.Now().Unix()-v.Time.Unix() < 60*60*24 {
			out.Count++
		}
	}
	context.JSON(http.StatusOK, out)
	logrus.Info("WebAPI: Response getAlarmPanels", context.Request.URL.Path+" done.")
}

// GetMachineInfo
// http://127.0.0.1:8088/api/devices/get-machine-info
// todo : Machine and Container
func GetMachineInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request getMachineInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	var out struct {
		Machine   int
		Container int
	}
	out.Machine = 0
	out.Container = 0
	context.JSON(http.StatusOK, out)
	logrus.Info("WebAPI: Response getMachineInfo", context.Request.URL.Path+" done.")
}
