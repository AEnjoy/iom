package rpc

import (
	"IOM/server/api"
	"IOM/server/config"
	"IOM/server/global"
	"IOM/server/model"
	pb "IOM/server/proto"
	"context"
	"errors"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type IOMProto struct{}

func (t *IOMProto) GetPackages(_ context.Context, packages *pb.Packages) (*pb.Flag, error) {
	var gPackages global.Packages
	var gPackage []global.Package
	gPackages.Token = packages.GetToken()
	gPackages.Type = packages.GetType()
	for _, p := range packages.GetPackages() {
		var dt global.Package
		dt.Arch = p.GetArch()
		dt.Description = p.GetDescription()
		dt.Version = p.GetVersion()
		dt.Name = p.GetName()
		dt.Repository = p.GetRepository()
		gPackage = append(gPackage, dt)
	}
	global.PackagesInfo[packages.GetToken()] = gPackages
	return &pb.Flag{Successful: true}, nil
}

func GTask2pTask(t *global.Task) *pb.Task {
	var resp pb.Task
	resp.Id = t.Id
	resp.Type = t.Type
	resp.Data = t.Args
	resp.HasScripts = t.HasScripts
	resp.ScriptsUrl = t.ScriptsUrl
	return &resp
}

func (t *IOMProto) GetTaskRequest(_ context.Context, request *pb.TaskRequest) (*pb.Task, error) {
	//
	ID := config.DeviceTokenGetDevicesID(request.GetToken())
	value, ok := global.TaskList[ID]
	if ok {
		delete(global.TaskList, ID) //任务派发完成
		return GTask2pTask(&value), nil
	}
	var dt pb.Task
	dt.Id = 0
	return &dt, errors.New("任务不存在或设备未连接")
}

var IOM IOMProto

func (t *IOMProto) GetSystemInfo(_ context.Context, r *pb.Host) (*pb.Flag, error) {
	var deviceT global.DevicesInfo
	var returnFlag pb.Flag
	clientID := r.Token
	if clientID == "" {
		return &pb.Flag{Successful: false, Tips: "token为空"}, errors.New("token为空")
	}
	host := model.PB2Host(r)
	dt, ok := global.DevicesInfos[clientID]
	if ok {
		deviceT = dt
	}
	deviceT.DataTime = time.Now()
	deviceT.HostInfo = host
	deviceT.FormatTime = deviceT.DataTime.Format("2006-01-02 15:04:05")
	global.DevicesInfos[clientID] = deviceT
	global.SetDeviceOnline(clientID)
	//debug
	//global.DevicesPrintDebugC(global.DevicesInfos[clientID])
	returnFlag.Successful = true
	return &returnFlag, nil
}

func (t *IOMProto) GetSystemState(_ context.Context, r *pb.State) (*pb.Flag, error) {
	var returnFlag pb.Flag
	var deviceT global.DevicesInfo
	clientID := r.Token
	if clientID == "" {
		return &pb.Flag{Successful: false, Tips: "token为空"}, errors.New("token为空")
	}
	state := model.PB2State(r)
	dt, ok := global.DevicesInfos[clientID]
	if ok {
		deviceT = dt
	}
	deviceT.DataTime = time.Now()
	deviceT.HostState = state
	deviceT.FormatTime = deviceT.DataTime.Format("2006-01-02 15:04:05")

	global.DevicesInfos[clientID] = deviceT
	global.SetDeviceOnline(clientID)
	//debug
	//global.DevicesPrintDebugC(global.DevicesInfos[clientID])
	//global.TrustedTokens[clientID] = true
	returnFlag.Successful = true
	return &returnFlag, nil
}

func ServerStart() {
	//for global.Port == "" {
	//	time.Sleep(3 * time.Second)
	//}
	listener, err := net.Listen("tcp", ":"+global.RPCPort)
	defer listener.Close()
	if err != nil {
		//logrus.Error("failed to listen: %v. The port is used.", err)
		api.Exit(1, "Failed to listen: %v. The port is used. Exit.")
	}
	s := grpc.NewServer()
	pb.RegisterIOMInfoServiceServer(s, &IOM)
	logrus.Info("Server listen in port:", global.RPCPort)
	err = s.Serve(listener)
}
