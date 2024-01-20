package console

import (
	"IOM/server/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

/*
Devices [[id] [[func] [args...]]]
DevicesAdd
DevicesDel [id]
DevicesShow [id]
DevicesList
Group [[id] [func] [args]]
GroupAdd
GroupDel
GroupShow
GroupList
*/

func Devices(Command []string) {

}

func DevicesAdd(Command []string) {
	var devicename, devicegroup, weightt string
	fmt.Print("新设备名>>> ")
	fmt.Scanln(&devicename)
	//devicename, _ := global.Read.ReadString('\n')
	fmt.Print("设备组名>>> ")
	fmt.Scanln(&devicegroup)
	//devicegroup, _ := global.Read.ReadString('\n')
	fmt.Print("设备权重[默认100]>>> ")
	fmt.Scanln(&weightt)
	//weightt, _ := global.Read.ReadString('\n')
	//weightt = strings.Replace(weightt, "\n", "", -1)
	var weight int
	if len(weightt) == 0 {
		weight = 100
	} else {
		weight, _ = strconv.Atoi(weightt)
	}

	var devicesID int
	devicesID, _ = captcha1()
	token := capthca2()
	flag := 0

	//devicename = strings.Replace(devicename, "\n", "", -1)
	//devicegroup = strings.Replace(devicegroup, "\n", "", -1)

	dt := config.GroupNameGetGroupID(devicegroup)
	if dt != 0 {
		config.DeviceAdd(devicesID, weight, token, devicename, flag, dt)
		println("设备添加成功\n设备Token为:", token)
		println("设备ID为:", devicesID)
		println("设备组ID为:", dt)
		logrus.Info("设备添加成功！设备Token为:", token)
	} else {
		fmt.Println("设备组名不存在！")
		return
	}

}
func DevicesDel(Command []string) {
	var devicesID int
	var gID int
	fmt.Print("设备ID>>> ")
	fmt.Scanln(&devicesID)
	gID, _ = config.DeviceIDGetGroupID(devicesID)
	if gID != 0 {
		config.DeviceDelete(devicesID, gID)
		println("设备删除成功！")
		logrus.Info("设备删除成功！")
		return
	}
	println("设备删除失败！")
	logrus.Errorln("设备删除失败！")
}
func GroupAddC() {
	fmt.Print("新建组名>>> ")
	var text string
	fmt.Scanln(&text)
	rows, _ := config.Db.Query("select GroupName from GroupsA where GroupName=?", text)
	if rows.Next() {
		fmt.Println("设备组名已存在")
		return
	}
	gid, _ := captcha1()
	//config.Db.Exec("insert into GroupsA(ID,GroupName) values (?,?)", strconv.Itoa(gid), text)
	config.GroupAdd(gid, text)
}
func GroupDelC() {
	fmt.Print("新建组名>>> ")
	var text string
	fmt.Scanln(&text)
	rows, _ := config.Db.Query("select GroupName from GroupsA where GroupName=?", text)
	if rows.Next() {
		id := config.GroupNameGetGroupID(text)
		config.GroupDelete(id)
		println("删除成功")
		logrus.Info("删除成功")
	} else {
		fmt.Println("设备组名不存在")
		return
	}
}
