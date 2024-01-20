package config

import (
	"IOM/server/global"
	"time"
)

type IOMSystem struct {
	Account  string //Admin
	Password string //Admin
}

// 设备
type Devices struct {
	Weight       int
	DevicesID    int
	Token        string //clientToken
	DevicesName  string
	flag         int //0:StandDevices,1:PVE,2:OpenStack,3:k8sHost
	AddedTime    time.Time
	LastDataTime time.Time

	//Data
	DevicesInfo global.DevicesInfo //nil
}

type DevicesGroup struct {
	GroupName     string
	DevicesListID []Devices
}

var DevicesGroupEX map[int]DevicesGroup //ID,DevicesGroupE
// 组
var SystemInfo []IOMSystem

func MapInit() {
	DevicesGroupEX = make(map[int]DevicesGroup)
	//DevicesEX = make(map[int]Devices)
}
