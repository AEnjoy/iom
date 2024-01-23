package api

import (
	"IOM/server/global"
)

func Main() {
	global.Router.POST("/api/group/add", groupAdd)
	global.Router.DELETE("/api/group/delete", groupDelete)
	global.Router.GET("/api/captcha", createCode) //pic
	global.Router.POST("/api/captcha/verify", verify)
	global.Router.POST("/api/user-password-change", passwordChange)
}
