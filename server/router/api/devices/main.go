package devices

import (
	"IOM/server/global"
)

func Main() {
	//global.Router.GET("/api/devices/add", add)
	global.Router.POST("/api/devices/add", add)
	global.Router.DELETE("/api/devices/delete", deleteDevice)
	global.Router.GET("/api/devices/get-devices", getDevices)
	global.Router.POST("/api/devices/get-devices", getDevices)
	global.Router.GET("/api/devices/get-all-devices", getAllDevices)
	global.Router.POST("/api/devices/get-all-devices", getAllDevices)
	global.Router.GET("/api/devices/get-panels", getPanels)
	global.Router.GET("/api/devices/get-alarm-panels", getAlarmPanels)
	global.Router.GET("/api/v2/devices/get-all-devices", getAllDevicesV2)
	global.Router.POST("/api/v2/devices/get-all-devices", getAllDevicesV2)
	global.Router.GET("/api/devices/:id/info", deviceInfo)
	global.Router.POST("/api/devices/:id/info", deviceInfo)
	global.Router.GET("/api/devices/:id/getpackages", getPackages)
	global.Router.POST("/api/devices/:id/getpackages", getPackages)
	global.Router.GET("/api/devices/:id/getdevices-name", getDevicesName)
}
