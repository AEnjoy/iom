package console

import (
	"IOM/server/api"
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
	devicesID, _ = api.Captcha1()
	token := api.Capthca2()
	flag := 0

	//devicename = strings.Replace(devicename, "\n", "", -1)
	//devicegroup = strings.Replace(devicegroup, "\n", "", -1)

	dt := config.GroupNameGetGroupID(devicegroup)
	if dt != 0 {
		config.DeviceAdd(devicesID, weight, token, devicename, flag, dt)
		println("设备添加成功\n设备Token为:", token)
		println("设备ID为:", devicesID)
		println("设备组ID为:", dt)
		logrus.Info("DevicesAdd:设备添加成功！设备Token为:", token)
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
		logrus.Info("DevicesDel:设备删除成功！ DevicesID:", devicesID)
		return
	}
	println("设备删除失败！")
	logrus.Errorln("E:DevicesDel:设备删除失败！")
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
	gid, _ := api.Captcha1()
	//config.Db.Exec("insert into GroupsA(ID,GroupName) values (?,?)", strconv.Itoa(gid), text)
	err := config.GroupAdd(gid, text)
	if err != nil {
		logrus.Errorln("E:GroupAddC:", err)
		return
	}
	logrus.Info("GroupAddC:", text, "成功！")
}
func GroupDelC() {
	fmt.Print("欲删除的组名>>> ")
	var text string
	fmt.Scanln(&text)
	rows, _ := config.Db.Query("select GroupName from GroupsA where GroupName=?", text)
	if rows.Next() {
		id := config.GroupNameGetGroupID(text)
		err := config.GroupDelete(id)
		if err != nil {
			logrus.Errorln("E:GroupDelC:", err)
			fmt.Println("删除失败")
			return
		}
		println("删除成功")
		logrus.Info("GroupDelC:删除成功")
	} else {
		fmt.Println("设备组名不存在")
		return
	}
}
