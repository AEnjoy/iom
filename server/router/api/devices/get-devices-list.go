package devices

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetDevices
// http://127.0.0.1:8088/api/devices/get-devices?token=xxx
func GetDevices(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	config.AddOperLogs("WebAPI访问：GetDevices", "GET|POST", context.Request.URL.Path)
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
		context.JSON(http.StatusOK, gin.H{})
		//context.String(http.StatusOK, "No Online devices")
		logrus.Info("WebAPI:  Request ", context.Request.URL.Path, " Get Online Devices successfully. But no online devices")
	}
}

// GetAllDevices 获取所有设备信息（包括不在线的） (不指定groupID则获取全部)
// http://127.0.0.1:8088/api/devices/get-all-devices?token=xxx[&groupID=xxx]
func GetAllDevices(context *gin.Context) {
	logrus.Info("WebAPI: Request getAllDevices: ", context.Request.URL.Path)
	config.AddOperLogs("WebAPI访问：GetAllDevices", "GET|POST", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	gID, _ := strconv.Atoi(context.DefaultQuery("groupID", "-1"))
	t1 := context.PostForm("groupID")
	if t1 != "" {
		gID, _ = strconv.Atoi(t1)
	}
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	type devicesList struct {
		Token string
		Type  int //0:StandDevices,1:PVE,2:OpenStack,3:k8sHost
	}
	var dt []devicesList
	if gID == -1 {
		for _, info := range config.DevicesGroupEX {
			for _, dev := range info.DevicesListID {
				var t devicesList
				t.Token = dev.Token
				t.Type = dev.Flag
				dt = append(dt, t)
			}
		}
	} else if config.IsGroupIDValid(gID) {
		for _, dev := range config.DevicesGroupEX[gID].DevicesListID {
			var t devicesList
			t.Token = dev.Token
			t.Type = dev.Flag
			dt = append(dt, t)
		}
	} else {
		context.String(http.StatusBadRequest, "groupID is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " groupID is invalid.")
		return
	}
	context.JSON(http.StatusOK, dt)
	logrus.Info("WebAPI: Request ", context.Request.URL.Path, " success.")
}

// GetAllDevicesV2 获取所有设备信息（包括不在线的） (不指定groupID则获取全部)
// http://127.0.0.1:8088/api/v2/devices/get-all-devices?token=xxx[&groupID=xxx]
func GetAllDevicesV2(context *gin.Context) {
	logrus.Info("WebAPI: Request getAllDevices: ", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	gID, _ := strconv.Atoi(context.DefaultQuery("groupID", "-1"))
	t1 := context.PostForm("groupID")
	if t1 != "" {
		gID, _ = strconv.Atoi(t1)
		config.AddOperLogs("WebAPI访问：GetAllDevicesV2", "POST", context.Request.URL.Path)
	} else {
		config.AddOperLogs("WebAPI访问：GetAllDevicesV2", "GET", context.Request.URL.Path)
	}
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	type devicesList struct {
		Name      string
		GroupID   int
		GroupName string
		Token     string
		Type      int //0:StandDevices,1:PVE,2:OpenStack,3:k8sHost
	}
	var dt []devicesList
	if gID == -1 {
		for _, info := range config.DevicesGroupEX {
			for _, dev := range info.DevicesListID {
				var t devicesList
				t.Name = dev.DevicesName
				t.GroupName = info.GroupName
				t.GroupID = config.GroupNameGetGroupID(info.GroupName)
				t.Token = dev.Token
				t.Type = dev.Flag
				dt = append(dt, t)
			}
		}
	} else if config.IsGroupIDValid(gID) {
		for _, dev := range config.DevicesGroupEX[gID].DevicesListID {
			var t devicesList
			t.Token = dev.Token
			t.Name = dev.DevicesName
			t.GroupID = gID
			t.GroupName = config.DevicesGroupEX[gID].GroupName
			t.Type = dev.Flag
			dt = append(dt, t)
		}
	} else {
		context.String(http.StatusBadRequest, "groupID is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " groupID is invalid.")
		return
	}
	context.JSON(http.StatusOK, dt)
	logrus.Info("WebAPI: Request ", context.Request.URL.Path, " success.")
}
