package router

import (
	"IOM/server/global"
	"IOM/server/router/api"
	"IOM/server/router/api/auth"
	"IOM/server/router/api/devices"
	"IOM/server/router/api/log"
	otherApis2 "IOM/server/router/api/otherApis"
	"IOM/server/router/api/res/file"
	"IOM/server/router/dashboard"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	global.Router = gin.Default()
	global.Router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			//Access-Control-Allow-Origin是必须的,他的值要么是请求Origin字段的值,要么是一个*, 表示接受任意域名的请求
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			//该字段可选。CORS请求时，XMLHttpRequest对象的getResponseHeader()方法只能拿到6个基本字段：Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma。
			//如果想拿到其他字段，就必须在Access-Control-Expose-Headers里面指定。上面的例子指定，getResponseHeader('FooBar')可以返回FooBar字段的值。
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, IOMAuth")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			//该字段可选，用来指定本次预检请求的有效期，单位为秒。有效期是20天（1728000秒），即允许缓存该条回应1728000秒（即20天），在此期间，不用发出另一条预检请求。
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				logrus.Panic("Panic info is: %v", err)
			}
		}()

		c.Next()
	})
	//路由注册
	regRouter()
	//静态资源路由处理
	static()
	// auth api:定时清除过期token
	go auth.TimeToClen()
	// server start
	err := global.Router.Run(":" + global.WebPort)
	if err != nil {
		logrus.Errorln("Web Server start error, error Info is: ", err)
		logrus.Info("Web Server will not start, please check the configuration or error Info.")
		return
	}
}

func regRouter() {
	//Auth Apis
	global.Router.GET("/api/auth/token/getToken", auth.TokenAuth)
	global.Router.GET("/api/auth/signin", auth.SignInAuth)
	global.Router.GET("/auth/signin", auth.SignInAuth)
	global.Router.POST("/auth/signin", auth.SignInAuthP)
	global.Router.POST("/api/auth/signin", auth.SignInAuthP)
	//Files Apis
	global.Router.GET("/api/res/file/:file", file.File)
	//Devices Apis
	global.Router.POST("/api/devices/add", devices.Add)
	global.Router.DELETE("/api/devices/delete", devices.DeleteDevice)
	global.Router.GET("/api/devices/get-devices", devices.GetDevices)
	global.Router.POST("/api/devices/get-devices", devices.GetDevices)
	global.Router.GET("/api/devices/get-all-devices", devices.GetAllDevices)
	global.Router.POST("/api/devices/get-all-devices", devices.GetAllDevices)
	global.Router.GET("/api/devices/get-panels", devices.GetPanels)
	global.Router.GET("/api/devices/get-alarm-panels", devices.GetAlarmPanels)
	global.Router.GET("/api/devices/get-machine-devices", devices.GetMachineInfo)
	global.Router.GET("/api/v2/devices/get-all-devices", devices.GetAllDevicesV2)
	global.Router.POST("/api/v2/devices/get-all-devices", devices.GetAllDevicesV2)
	global.Router.GET("/api/devices/:id/info", devices.DeviceInfo)
	global.Router.POST("/api/devices/:id/info", devices.DeviceInfo)
	global.Router.GET("/api/devices/:id/getpackages", devices.GetPackages)
	global.Router.POST("/api/devices/:id/getpackages", devices.GetPackages)
	global.Router.GET("/api/devices/:id/getdevices-name", devices.GetDevicesName)
	//Group Apis
	global.Router.POST("/api/group/add", api.GroupAdd)
	global.Router.DELETE("/api/group/delete", api.GroupDelete)
	global.Router.GET("/api/group/search/name-get-id", api.GroupNameGetGroupIDGetMethod)
	global.Router.GET("/api/group/id/is-valid", api.IsGroupIDValidGetMethod)
	global.Router.GET("/api/group/get-list", api.GetGroupList)
	global.Router.POST("/api/group/search/name-get-id", api.GroupNameGetGroupIDPostMethod)
	global.Router.POST("/api/group/id/is-valid", api.IsGroupIDValidPostMethod)
	//Verify Apis
	global.Router.GET("/api/captcha", api.CreateCode) //pic
	global.Router.POST("/api/captcha/verify", api.Verify)
	global.Router.POST("/api/user-password-change", api.PasswordChange)
	//DashBoard Apis
	global.Router.GET("/api/dashboard/user/info", dashboard.GetUserInfo)
	//LogsSystem Apis
	global.Router.GET("/api/log/logLogin/list", log.ListLoginInfo)
	global.Router.DELETE("/api/log/logLogin/remove/:infoID", log.DelLoginInfo)
	global.Router.DELETE("/api/log/logLogin/clean-all", log.CleanAllLoginLog)
	global.Router.GET("/api/log/logOper/list", log.ListOperInfo)
	global.Router.DELETE("/api/log/logOper/remove/:id", log.DelOperInfo)
	global.Router.DELETE("/api/log/logOper/all", log.CleanAllOpernfo)
	//Other Apis
	global.Router.GET("/api/creat-token", otherApis2.CreatToken)
	global.Router.GET("/api/version", func(context *gin.Context) {
		context.String(200, global.Version)
	})
	//time.Sleep(time.Second / 2)
}
