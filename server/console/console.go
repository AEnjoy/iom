package console

import (
	"IOM/server/api"
	"IOM/server/config"
	"IOM/server/global"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var failTimes = 0

func Home() {
	goto login
login:
	//global.Read = bufio.NewReader(os.Stdin)
	var user config.IOMSystem
	fmt.Print("UserName>>> ")
	fmt.Scanln(&user.Account)
	if user.Account == "" { //有时候会因为输入回车清屏，导致user.Account为空
		goto login
	}
	fmt.Print("PassWord>>> ")
	fmt.Scanln(&user.Password)
	if config.IsUserInTrustUsers(user) {
		fmt.Println("Login Success!")
		fmt.Println("---------------------------------------------------------------------------")
		fmt.Println("Type \"version\" to show the IOMServer Version Information.")
		fmt.Println("Type \"help\" to get the help messages(and all command).")
		fmt.Println("Type \"exit\" or \"quit\" to exit.")
		fmt.Println("---------------------------------------------------------------------------")
	} else if !config.IsUserInTrustUsers(user) && failTimes <= 3 {
		failTimes++
		fmt.Println("Login Failed!\nPlease Check Your UserName and PassWord.\n" +
			"If you have not remembered your PassWord, please \n" +
			"run \"IOMServer --getAdmin\" on the server to obtain the administrator account\n" +
			" and password, and perform the reset operation after the login is successful.")
		goto login
	} else if failTimes > 3 {
		fmt.Println("登录失败过多，请三分钟后再试.")
		logrus.Info("Console：登录失败")
		time.Sleep(3 * time.Minute)
		failTimes = 0
		goto login
	}
	time.Sleep(1000 * time.Millisecond)
	for {
		fmt.Print(">>> ")
		var text string
		fmt.Scanln(&text)
		//text, _ := global.Read.ReadString('\n')
		// convert CRLF to LF
		// Win: text = strings.Replace(text,"\r\n","",-1)
		text = strings.Replace(text, "\n", "", -1)
		s := strings.Split(text, " ")
		if len(s) == 0 {
			continue
		}
		if strings.Compare("DevicesAdd", s[0]) == 0 {
			DevicesAdd(s[1:])
			continue
		}
		if strings.Compare("DevicesDel", s[0]) == 0 {
			DevicesDel(s[1:])
			continue
		}
		if strings.Compare("GroupAdd", s[0]) == 0 {
			GroupAddC()
			continue
		}
		if strings.Compare("GroupDel", s[0]) == 0 {
			GroupDelC()
			continue
		}
		if strings.Compare("quit", s[0]) == 0 {
			config.DBExitSave()
			api.Exit(0, "Command quit")
		}
		if strings.Compare("exit", s[0]) == 0 {
			config.DBExitSave()
			api.Exit(0, "Command exit")
		}
		if strings.Compare("help", s[0]) == 0 {
			fmt.Printf(global.Help, global.Version)
			continue
		}
		if strings.Compare("version", s[0]) == 0 {
			fmt.Println(global.Version)
			continue
		}
		if strings.Compare("ShowDevicesState", s[0]) == 0 {
			//PC-test
			//94275941 非必须
			//47063971
			fmt.Print("设备组ID>>> ")
			fmt.Scanln(&text)
			idG, _ := strconv.Atoi(text)
			fmt.Print("设备ID>>> ")
			fmt.Scanln(&text)
			idD, _ := strconv.Atoi(text)
			DashboardShowDevicesStateC(idG, idD)
			continue
		}
		if strings.Compare("ShowDevicesPackages", s[0]) == 0 {
			fmt.Print("设备ID或设备Token>>> ")
			fmt.Scanln(&text)
			idD, e := strconv.Atoi(text)
			if e != nil {
				//token
				PrintPackagesC(config.DeviceTokenGetDevicesID(text))
			} else {
				PrintPackagesC(idD)
			}
			continue
		}
		if strings.Compare("ConfigEdit", s[0]) == 0 {
			ConfigEditC()
			continue
		}
		if strings.Compare("AdminAdd", s[0]) == 0 {
			AdminAddC()
			continue
		}
		if strings.Compare("AdminDel", s[0]) == 0 {
			AdminDelC()
			continue
		}
		if strings.Compare("AdminEdit", s[0]) == 0 {
			AdminEditC()
			continue
		}
		fmt.Println("Unknown command:", s[0])
	}
}
