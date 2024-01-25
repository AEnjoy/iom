package api

import (
	"IOM/server/global"
)

func Main() {
	global.Router.POST("/api/group/add", groupAdd)
	global.Router.DELETE("/api/group/delete", groupDelete)
	global.Router.GET("/api/captcha", createCode) //pic
	global.Router.GET("/api/group/search/name-get-id", groupNameGetGroupIDGetMethod)
	global.Router.GET("/api/group/id/is-valid", isGroupIDValidGetMethod)
	global.Router.POST("/api/captcha/verify", verify)
	global.Router.POST("/api/user-password-change", passwordChange)
	global.Router.POST("/api/group/search/name-get-id", groupNameGetGroupIDPostMethod)
	global.Router.POST("/api/group/id/is-valid", isGroupIDValidPostMethod)
}
