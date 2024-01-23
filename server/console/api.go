package console

import (
	"IOM/server/config"
	"IOM/server/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

func DashboardShowDevicesStateC(groupID int, devicesID int) {
	token := config.DeviceIDGetDeviceToken(devicesID)
	if token != "" {
		dt := global.DevicesInfos[token]
		println("\t----\t")
		println("设备组ID:", groupID)
		println("设备ID:", devicesID)
		println("数据报时间:", dt.DataTime.Format("2006-01-02 15:04:05"))
		println("客户端ID(Token):", dt.ClientID)
		println("内存占用:", dt.HostState.MemUsed/1024/1024, "MB")
		println("CPU占用:", dt.HostState.CPU, "%")
		println("磁盘占用:", dt.HostState.DiskUsed/1024/1024, "MB")
		println("网络下行速度:", dt.HostState.NetInSpeed/1024/1024, "MB")
		println("网络上行速度:", dt.HostState.NetOutSpeed/1024/1024, "MB")
		println("CPULoad(1,5,15):", dt.HostState.Load1, ",", dt.HostState.Load5, ",", dt.HostState.Load15)
		println("TCP,UDP连接量:", dt.HostState.TcpConnCount, ",", dt.HostState.UdpConnCount)
		println("系统进程数:", dt.HostState.ProcessCount)
		println("\t----\t")
		println("系统信息:", dt.HostInfo.Platform, " ", dt.HostInfo.PlatformVersion)
		println("系统架构:", dt.HostInfo.Arch)
		println("总内存:", dt.HostInfo.MemTotal)
		println("硬盘总大小:", dt.HostInfo.DiskTotal)
		println("交换分区大小:", dt.HostInfo.SwapTotal)
		println("系统启动时间:", dt.HostInfo.BootTime)
		println("IP地址:", dt.HostInfo.IP)
		println("国家代码:", dt.HostInfo.CountryCode)
		println("\t----\t")
		return
	}
	logrus.Error("E:未找到设备信息")
}
func ConfigEditC() {
	var confName string
	var confValue string
	fmt.Println("配置项名:")
	fmt.Scanln(&confName)
	rows, _ := config.Db.Query("select * from Config where ConfigName = ?", confName)
	defer rows.Close()
	if rows.Next() {
		fmt.Println("配置值值:")
		fmt.Scanln(&confValue)
		config.Edit(-1, confName, confValue)
		logrus.Info("配置项修改-name，value:", confName, confValue)
	} else {
		logrus.Errorln("E:配置项不存在")
	}
}
func AdminAddC() {
	var adminName string
	var adminPwd string
	fmt.Println("管理员账号:")
	fmt.Scanln(&adminName)
	fmt.Println("管理员密码:")
	fmt.Scanln(&adminPwd)
	config.AdminAccountAdd(adminName, adminPwd)
	logrus.Info("管理员添加-name，pwd:", adminName, adminPwd)
}
func AdminDelC() {
	var adminName string
	fmt.Println("管理员账号:")
	fmt.Scanln(&adminName)
	config.AdminAccountDel(adminName)
	logrus.Info("管理员删除-name:", adminName)
}
func AdminEditC() {
	var adminName string
	var adminPwd string
	fmt.Println("管理员账号:")
	fmt.Scanln(&adminName)
	fmt.Println("新密码:")
	fmt.Scanln(&adminPwd)
	config.AdminAccountEdit(adminName, adminPwd)
	logrus.Info("管理员密码修改-name，pwd:", adminName, adminPwd)
}
func PrintPackagesC(id int) {
	token := config.DeviceIDGetDeviceToken(id)
	dt, ok := global.PackagesInfo[token]
	if ok {
		logrus.Info("软件包信息输出Console")
		println("DevicesToken:", dt.Token, ' ', "DevicesID:", id)
		for _, p := range dt.Packages {
			println("软件包名:", p.Name, " 软件包版本:", p.Version, " 软件包架构:", p.Arch, " 软件包存储库:", p.Repository, " 软件包描述信息:", p.Description)
		}
	}
	logrus.Error("找不到该设备的软件包信息")
}
