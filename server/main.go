package main

import (
	IOM "IOM/server/api"
	"IOM/server/config"
	"IOM/server/console"
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
)

// 自定义初始化
func Init() {
	//log init
	logf, _ := os.OpenFile("./logs/iomServer.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	mw := io.MultiWriter(os.Stdout, logf)
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	go IOM.ExitHandle(exitChan)
	logrus.SetOutput(mw)
	//防止重复运行 lock init
	go IOM.Lock()

	go config.MapInit()
	go global.MapInit()
	global.TaskListInit()
	//go urltools.Handle()
	config.DBOpen()
	go rpc.ServerStart()
	go global.TimeToCheckOnline()
}

func main() {
	var shell = true
	var server = true
	var admin = true
	var version = true
	//flagSet := pflag.NewFlagSet("server", pflag.ContinueOnError)
	pflag.BoolVarP(&shell, "shell", "c", false, "启动 IO&M console管理控制台")
	pflag.BoolVarP(&server, "web", "w", false, "启动 IO&M Web管理面板")
	pflag.BoolVar(&admin, "getAdmin", false, "Show IO&M Dashboard Admin Password and Username and exit.")
	pflag.BoolVarP(&version, "version", "V", false, "Show IO&MServer version and exit.")
	pflag.Parse()

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
	if server == false && shell == false {
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
Welcome to IO&M Console Version:%s 
\n\nPlease Login:\n
`, global.Version)
		go console.Home()
	}

	select {}
}
