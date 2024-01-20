package model

import (
	"IOM/server/global"
	pb "IOM/server/proto"
)

func (s *HostState) PB() *pb.State {
	return &pb.State{
		Cpu:            s.CPU,
		MemUsed:        s.MemUsed,
		SwapUsed:       s.SwapUsed,
		DiskUsed:       s.DiskUsed,
		NetInTransfer:  s.NetInTransfer,
		NetOutTransfer: s.NetOutTransfer,
		NetInSpeed:     s.NetInSpeed,
		NetOutSpeed:    s.NetOutSpeed,
		Uptime:         s.Uptime,
		Load1:          s.Load1,
		Load5:          s.Load5,
		Load15:         s.Load15,
		TcpConnCount:   s.TcpConnCount,
		UdpConnCount:   s.UdpConnCount,
		ProcessCount:   s.ProcessCount,
	}
}
func PB2State(s *pb.State) global.HostState {
	return global.HostState{
		CPU:            s.GetCpu(),
		MemUsed:        s.GetMemUsed(),
		SwapUsed:       s.GetSwapUsed(),
		DiskUsed:       s.GetDiskUsed(),
		NetInTransfer:  s.GetNetInTransfer(),
		NetOutTransfer: s.GetNetOutTransfer(),
		NetInSpeed:     s.GetNetInSpeed(),
		NetOutSpeed:    s.GetNetOutSpeed(),
		Uptime:         s.GetUptime(),
		Load1:          s.GetLoad1(),
		Load5:          s.GetLoad5(),
		Load15:         s.GetLoad15(),
		TcpConnCount:   s.GetTcpConnCount(),
		UdpConnCount:   s.GetUdpConnCount(),
		ProcessCount:   s.GetProcessCount(),
	}
}
func (h *Host) PB() *pb.Host {
	return &pb.Host{
		Platform:        h.Platform,
		PlatformVersion: h.PlatformVersion,
		Cpu:             h.CPU,
		MemTotal:        h.MemTotal,
		DiskTotal:       h.DiskTotal,
		SwapTotal:       h.SwapTotal,
		Arch:            h.Arch,
		Virtualization:  h.Virtualization,
		BootTime:        h.BootTime,
		Ip:              h.IP,
		CountryCode:     h.CountryCode,
		Version:         h.Version,
	}
}
func PB2Host(h *pb.Host) global.Host {
	return global.Host{
		Platform:        h.GetPlatform(),
		PlatformVersion: h.GetPlatformVersion(),
		CPU:             h.GetCpu(),
		MemTotal:        h.GetMemTotal(),
		DiskTotal:       h.GetDiskTotal(),
		SwapTotal:       h.GetSwapTotal(),
		Arch:            h.GetArch(),
		Virtualization:  h.GetVirtualization(),
		BootTime:        h.GetBootTime(),
		IP:              h.GetIp(),
		CountryCode:     h.GetCountryCode(),
		Version:         h.GetVersion(),
	}
}

// there are some bugs that the code has to be modified
func HostState2DevicesInfo(rs global.HostState, clineID string, dt *global.DevicesInfo) {
	if dt != nil && dt.ClientID == clineID {
		dt.HostState = rs
	} else {
		var dtv global.DevicesInfo
		dtv.ClientID = clineID
		dtv.HostState = rs
		dt = &dtv
	}
}
func Host2DevicesInfo(rs global.Host, clineID string, dt *global.DevicesInfo) {
	if dt != nil && dt.ClientID == clineID {
		dt.HostInfo = rs
	} else {
		var dtv global.DevicesInfo
		dtv.ClientID = clineID
		dtv.HostInfo = rs
		dt = &dtv
	}
}
func ClientIDGetDevicesInfo(clineID string, dt *global.DevicesInfo) {
	dT, ok := global.DevicesInfos[clineID]
	//println("Debug:ok", ok)
	if ok {
		dt = &dT
		return
	}
	dt = nil
}
