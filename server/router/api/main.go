package api

import (
	"IOM/server/global"
	"github.com/gin-gonic/gin"
)

func Main() {
	global.Router.POST("/api/group/add", groupAdd)
	global.Router.DELETE("/api/group/delete", groupDelete)
	global.Router.GET("/api/version", func(context *gin.Context) {
		context.String(200, global.Version)
	})
	global.Router.GET("/api/captcha", createCode) //pic
	global.Router.POST("/api/captcha/verify", verify)
	global.Router.POST("/api/user-password-change", passwordChange)
}
