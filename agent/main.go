package main

import (
	"IOM/agent/model"
	"IOM/agent/pkg/monitor"
	"IOM/agent/pkg/pkgmgr"
	"IOM/agent/pkg/util"
	pb "IOM/agent/proto"
	"context"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	bpc "github.com/DaRealFreak/cloudflare-bp-go"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	//"github.com/nezhahq/go-github-selfupdate/selfupdate"
	"github.com/shirou/gopsutil/v3/disk"
	psnet "github.com/shirou/gopsutil/v3/net"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
)

//-X main.version=0.1alpha -X main.arch=amd64

type AgentCliParam struct {
	SkipConnectionCount   bool   // 跳过连接数检查
	SkipProcsCount        bool   // 跳过进程数量检查
	DisableAutoUpdate     bool   // 关闭自动更新
	DisableForceUpdate    bool   // 关闭强制更新
	DisableCommandExecute bool   // 关闭命令执行
	Debug                 bool   // debug模式
	Server                string // 服务器地址
	ClientSecret          string // 客户端密钥（Token）
	ReportDelay           int    // 报告间隔
	TLS                   bool   // 是否使用TLS加密传输至服务端
	Version               bool   // 当前版本号

	RunMode int //0:[default] 1:PVE envTool 2:k8s envTool
}

var (
	version string
	arch    string
	client  pb.IOMInfoServiceClient
	inited  bool
)

var (
	agentCliParam AgentCliParam
	agentConfig   model.AgentConfig
	httpClient    = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 30,
	}
)

const (
	delayWhenError = time.Second * 10 // Agent 重连间隔
	networkTimeOut = time.Second * 5  // 普通网络超时
	macOSChromeUA  = ""
)

func init() {
	net.DefaultResolver.PreferGo = true // 使用 Go 内置的 DNS 解析器解析域名
	net.DefaultResolver.Dial = func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{
			Timeout: time.Second * 5,
		}
		dnsServers := util.DNSServersAll
		if len(agentConfig.DNS) > 0 {
			dnsServers = agentConfig.DNS
		}
		index := int(time.Now().Unix()) % int(len(dnsServers))
		queue := generateQueue(index, len(dnsServers))
		var conn net.Conn
		var err error
		for i := 0; i < len(queue); i++ {
			conn, err = d.DialContext(ctx, "udp", dnsServers[queue[i]])
			if err == nil {
				return conn, nil
			}
		}
		return nil, err
	}
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true

	http.DefaultClient.Timeout = time.Second * 5
	httpClient.Transport = bpc.AddCloudFlareByPass(httpClient.Transport, bpc.Options{
		AddMissingHeaders: true,
		Headers: map[string]string{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			"Accept-Language": "en-US,en;q=0.5",
			"User-Agent":      monitor.MacOSChromeUA,
		},
	})

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	agentConfig.Read(filepath.Dir(ex) + "/config.yml")
}

func main() {
	// windows环境处理

	arch = "amd64"
	version = "0.0.1"

	/*if runtime.GOOS == "windows" {
		hostArch, err := host.KernelArch()
		if err != nil {
			panic(err)
		}
		if hostArch == "i386" {
			hostArch = "386"
		}
		if hostArch == "i686" || hostArch == "ia64" || hostArch == "x86_64" {
			hostArch = "amd64"
		}
		if hostArch == "aarch64" {
			hostArch = "arm64"
		}
		if arch != hostArch {
			panic(fmt.Sprintf("与当前系统不匹配，当前运行 %s_%s, 需要下载 %s_%s", runtime.GOOS, arch, runtime.GOOS, hostArch))
		}
	}*/

	// 来自于 GoReleaser 的版本号
	monitor.Version = version

	// 初始化运行参数
	var isEditAgentConfig bool
	flag.BoolVarP(&agentCliParam.Debug, "debug", "d", false, "开启调试信息")
	flag.BoolVarP(&isEditAgentConfig, "edit-agent-config", "c", false, "修改要监控的网卡/分区名单，修改自定义 DNS")
	flag.StringVarP(&agentCliParam.Server, "server", "s", "localhost:10000", "管理面板RPC端口")
	flag.StringVarP(&agentCliParam.ClientSecret, "password", "p", "", "Agent连接Secret")
	flag.IntVar(&agentCliParam.ReportDelay, "report-delay", 1, "系统状态上报间隔")
	flag.IntVar(&agentCliParam.RunMode, "run-mode", 0, "Agent运行模式 0：默认  1：PVE 2：k8sNode")
	flag.BoolVar(&agentCliParam.SkipConnectionCount, "skip-conn", false, "不监控连接数")
	flag.BoolVar(&agentCliParam.SkipProcsCount, "skip-procs", false, "不监控进程数")
	flag.BoolVarP(&agentCliParam.Version, "version", "v", false, "查看当前版本号")
	agentCliParam.DisableCommandExecute = false
	agentCliParam.DisableAutoUpdate = false
	agentCliParam.DisableForceUpdate = false
	agentCliParam.TLS = false
	//flag.BoolVar(&agentCliParam.DisableCommandExecute, "disable-command-execute", false, "禁止在此机器上执行命令")
	//flag.BoolVar(&agentCliParam.DisableAutoUpdate, "disable-auto-update", false, "禁用自动升级")
	//flag.BoolVar(&agentCliParam.DisableForceUpdate, "disable-force-update", false, "禁用强制升级")
	//flag.BoolVar(&agentCliParam.TLS, "tls", false, "启用SSL/TLS加密")
	flag.Parse()

	if agentCliParam.Version {
		fmt.Println(version)
		return
	}

	if isEditAgentConfig {
		editAgentConfig()
		return
	}

	if agentCliParam.ClientSecret == "" {
		flag.Usage()
		return
	}

	if agentCliParam.ReportDelay < 1 || agentCliParam.ReportDelay > 4 {
		println("report-delay 的区间为 1-4")
		return
	}

	run()
}

func run() {
	println("IOM Agent version:", version)
	// 下载远程命令执行需要的终端
	if !agentCliParam.DisableCommandExecute {
		//go pty.DownloadDependency()
	}
	// 上报服务器信息
	go reportState()
	// 更新IP信息
	go monitor.UpdateIP()
	go timeToGetTask()
	// 定时检查更新
	/*
		if _, err := semver.Parse(version); err == nil && !agentCliParam.DisableAutoUpdate {
			doSelfUpdate(true)
			go func() {
				for range time.Tick(20 * time.Minute) {
					doSelfUpdate(true)
				}
			}()
		}
	*/
	//上报软件包信息
	go reportPackage()

	var err error
	var conn *grpc.ClientConn

	retry := func() {
		inited = false
		println("Error to close connection ...")
		if conn != nil {
			conn.Close()
		}
		time.Sleep(delayWhenError)
		println("Try to reconnect ...")
	}

	for {
		timeOutCtx, cancel := context.WithTimeout(context.Background(), networkTimeOut)
		time.Sleep(5 * time.Millisecond)
		conn, err = grpc.DialContext(timeOutCtx, agentCliParam.Server, grpc.WithInsecure())
		if err != nil {
			println("与面板建立连接失败：", err)
			cancel()
			retry()
			continue
		}
		cancel()
		client = pb.NewIOMInfoServiceClient(conn)
		timeOutCtx, cancel = context.WithTimeout(context.Background(), networkTimeOut)
		fmt.Println("上报系统信息...")
		pbT := monitor.GetHost(&agentConfig).PB()
		pbT.Token = agentCliParam.ClientSecret
		_, err = client.GetSystemInfo(timeOutCtx, pbT)
		if err != nil {
			println("上报系统信息失败：", err)
			cancel()
			retry()
			continue
		}
		time.Sleep(time.Second * 3)
		cancel()
		inited = true

		//retry()
	}
}

func timeToGetTask() {
	for {
		time.Sleep(time.Second * 10)
		task, _ := client.GetTaskRequest(context.Background(), &pb.TaskRequest{Token: agentCliParam.ClientSecret})
		if task != nil && task.Id != 0 {
			//有任务
		}
	}
}
func reportPackage() {
	var lastReportHostInfo time.Time
	var err error
	defer println("reportPackage exit", time.Now(), "=>", err)
	for {
		fmt.Println("软件包数据上报", time.Now())
		if lastReportHostInfo.Before(time.Now().Add(-10 * time.Minute)) {
			lastReportHostInfo = time.Now()
			packages := pkgmgr.TimeToGetPackage()
			if packages == nil {
				break
			}
			packages.Token = agentCliParam.ClientSecret
			client.GetPackages(context.Background(), packages)
		}
		time.Sleep(time.Second * 10)
	}
}

/*
func receiveTasks(tasks pb.NezhaService_RequestTaskClient) error {
	var err error
	defer println("receiveTasks exit", time.Now(), "=>", err)
	for {
		var task *pb.Task
		task, err = tasks.Recv()
		if err != nil {
			return err
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					println("task panic", task, err)
				}
			}()
			doTask(task)
		}()
	}
}

func doTask(task *pb.Task) {
	var result pb.TaskResult
	result.Id = task.GetId()
	result.Type = task.GetType()
	switch task.GetType() {
	case model.TaskTypeTerminal:
		handleTerminalTask(task)
	case model.TaskTypeHTTPGET:
		handleHttpGetTask(task, &result)
	case model.TaskTypeICMPPing:
		handleIcmpPingTask(task, &result)
	case model.TaskTypeTCPPing:
		handleTcpPingTask(task, &result)
	case model.TaskTypeCommand:
		handleCommandTask(task, &result)
	case model.TaskTypeUpgrade:
		handleUpgradeTask(task, &result)
	case model.TaskTypeKeepalive:
		return
	default:
		println("不支持的任务：", task)
	}
	client.ReportTask(context.Background(), &result)
}
*/
// reportState 向server上报状态信息
func reportState() {
	fmt.Println("数据上报", time.Now())
	var lastReportHostInfo time.Time
	var err error
	defer println("reportState exit", time.Now(), "=>", err)
	for {
		//time.Sleep(5 * time.Millisecond)
		// 为了更准确的记录时段流量，inited 后再上传状态信息
		if client != nil && inited {
			monitor.TrackNetworkSpeed(&agentConfig)
			timeOutCtx, cancel := context.WithTimeout(context.Background(), networkTimeOut)
			pbT := monitor.GetState(&agentConfig, agentCliParam.SkipConnectionCount, agentCliParam.SkipProcsCount).PB()
			pbT.Token = agentCliParam.ClientSecret
			_, err = client.GetSystemState(timeOutCtx, pbT)
			cancel()
			if err != nil {
				println("reportState error", err)
				time.Sleep(delayWhenError)
			}
			// 每10分钟重新获取一次硬件信息
			if lastReportHostInfo.Before(time.Now().Add(-10 * time.Minute)) {
				lastReportHostInfo = time.Now()
				pbT2 := monitor.GetHost(&agentConfig).PB()
				pbT2.Token = agentCliParam.ClientSecret
				client.GetSystemInfo(context.Background(), pbT2)
			}
		}
		time.Sleep(time.Second * time.Duration(agentCliParam.ReportDelay))
	}
}

// 修改Agent要监控的网卡与硬盘分区
func editAgentConfig() {
	nc, err := psnet.IOCounters(true)
	if err != nil {
		panic(err)
	}
	var nicAllowlistOptions []string
	for _, v := range nc {
		nicAllowlistOptions = append(nicAllowlistOptions, v.Name)
	}

	var diskAllowlistOptions []string
	diskList, err := disk.Partitions(false)
	if err != nil {
		panic(err)
	}
	for _, p := range diskList {
		diskAllowlistOptions = append(diskAllowlistOptions, fmt.Sprintf("%s\t%s\t%s", p.Mountpoint, p.Fstype, p.Device))
	}

	var qs = []*survey.Question{
		{
			Name: "nic",
			Prompt: &survey.MultiSelect{
				Message: "选择要监控的网卡",
				Options: nicAllowlistOptions,
			},
		},
		{
			Name: "disk",
			Prompt: &survey.MultiSelect{
				Message: "选择要监控的硬盘分区",
				Options: diskAllowlistOptions,
			},
		},
		{
			Name: "dns",
			Prompt: &survey.Input{
				Message: "自定义 DNS，可输入空格跳过，如 1.1.1.1:53,1.0.0.1:53",
				Default: strings.Join(agentConfig.DNS, ","),
			},
		},
	}

	answers := struct {
		Nic  []string
		Disk []string
		DNS  string
	}{}

	err = survey.Ask(qs, &answers, survey.WithValidator(survey.Required))
	if err != nil {
		fmt.Println("选择错误", err.Error())
		return
	}

	agentConfig.HardDrivePartitionAllowlist = []string{}
	for _, v := range answers.Disk {
		agentConfig.HardDrivePartitionAllowlist = append(agentConfig.HardDrivePartitionAllowlist, strings.Split(v, "\t")[0])
	}

	agentConfig.NICAllowlist = make(map[string]bool)
	for _, v := range answers.Nic {
		agentConfig.NICAllowlist[v] = true
	}

	dnsServers := strings.TrimSpace(answers.DNS)

	if dnsServers != "" {
		agentConfig.DNS = strings.Split(dnsServers, ",")
		for _, s := range agentConfig.DNS {
			host, _, err := net.SplitHostPort(s)
			if err == nil {
				if net.ParseIP(host) == nil {
					err = errors.New("格式错误")
				}
			}
			if err != nil {
				panic(fmt.Sprintf("自定义 DNS 格式错误：%s %v", s, err))
			}
		}
	} else {
		agentConfig.DNS = []string{}
	}

	if err = agentConfig.Save(); err != nil {
		panic(err)
	}

	fmt.Println("修改自定义配置成功，重启 Agent 后生效")
}

func println(v ...interface{}) {
	if agentCliParam.Debug {
		fmt.Printf("IOM@%s>> ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(v...)
	}
}

func generateQueue(start int, size int) []int {
	var result []int
	for i := start; i < start+size; i++ {
		if i < size {
			result = append(result, i)
		} else {
			result = append(result, i-size)
		}
	}
	return result
}
