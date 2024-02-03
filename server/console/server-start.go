package console

import (
	IOM "IOM/server/api"
	"IOM/server/config"
	"IOM/server/global"
	"IOM/server/router"
	"IOM/server/service/rpc"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 自定义初始化
func Init() {
	//log init
	//config.StartTime = time.Now()
	t := config.StartTime.Format("2006-01-02")
	logf, _ := os.OpenFile("./logs/"+t+"iomServer.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	mw := io.MultiWriter(os.Stdout, logf)
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	go IOM.ExitHandle(exitChan)
	logrus.SetOutput(mw)
	//防止重复运行 lock init
	IOM.Lock()
	logrus.Info("IOM server start. Time:" + time.Now().Format("2006-01-02 15:04:05"))
	go config.MapInit()
	go global.MapInit()
	global.TaskListInit()
	//go urltools.Handle()
	config.DBOpen()
	go config.LoadOnceToken()
	go rpc.ServerStart()
	go global.TimeToCheckOnline()
	go global.TimeToCheckTokenIsValid()
}
func Main() {
	var shell = true
	var server = true
	var admin = true
	var version = true
	//flagSet := pflag.NewFlagSet("server", pflag.ContinueOnError)
	pflag.BoolVarP(&shell, "shell", "c", false, "启动 IO&M console管理控制台")
	pflag.BoolVarP(&server, "web", "w", false, "启动 IO&M Web管理面板")
	pflag.BoolVar(&admin, "getAdmin", false, "Show IO&M Dashboard Admin Password and Username and exit.")
	pflag.BoolVarP(&global.DebugFlag, "debug", "d", false, "Debug mode 将显示一些额外信息，并忽略一些错误，可能会泄露某些数据。")
	pflag.BoolVarP(&version, "version", "V", false, "Show IO&MServer version and exit.")
	pflag.Parse()
	goDebugEnv := os.Getenv("GO_DEBUG")
	if goDebugEnv == "1" {
		global.DebugFlag = true
	}

	if admin {
		go config.MapInit()
		go global.MapInit()
		global.TaskListInit()
		config.DBOpen()
		rows, _ := config.Db.Query("select * from IOMSystem")
		for rows.Next() {
			var um, pw string
			rows.Scan(&um, &pw)
			println("可以登录的用户名：" + um + " 密码：" + pw)
		}
		config.DBExitSave()
		IOM.Exit(0, "")
	}
	if version {
		fmt.Println("IOM Server Version: ", global.Version)
		IOM.Exit(0, "")
	}
	if !server && !shell {
		pflag.Usage()
		IOM.Exit(1, "--server or --shell must be true (两项服务可同时开启)")
	}
	//Init
	Init()
	if server {
		go router.RouterInit()
	}
	if shell {
		fmt.Printf(`
Welcome to IO&M Console(Server) Version: %s 
Please Login:
`, global.Version)
		go home()
	}

	select {}
}
