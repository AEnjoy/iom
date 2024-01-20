package global

import "time"

type DevicesInfo struct {
	ClientID  string    //key
	HostState HostState //周期数据
	HostInfo  Host
	DataTime  time.Time
	Online    bool

	// 格式化时间
	FormatTime string
}

type HostState struct {
	CPU            float64
	MemUsed        uint64
	SwapUsed       uint64
	DiskUsed       uint64
	NetInTransfer  uint64
	NetOutTransfer uint64
	NetInSpeed     uint64
	NetOutSpeed    uint64
	Uptime         uint64
	Load1          float64
	Load5          float64
	Load15         float64
	TcpConnCount   uint64
	UdpConnCount   uint64
	ProcessCount   uint64
}

type Host struct {
	Platform        string
	PlatformVersion string
	CPU             []string
	MemTotal        uint64
	DiskTotal       uint64
	SwapTotal       uint64
	Arch            string
	Virtualization  string
	BootTime        uint64
	IP              string `json:"-"`
	CountryCode     string
	Version         string
}

// DevicesPrintDebugC  for debug
func DevicesPrintDebugC(dt DevicesInfo) {
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
}
