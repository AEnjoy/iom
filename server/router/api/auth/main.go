package auth

import (
	"IOM/server/config"
	"IOM/server/global"
)

var tokenApis []*config.TokenApi //know

// Main  http://127.0.0.1:8088/api/auth/token/whatever
func Main() {
	//http://127.0.0.1:8088/api/auth/token/getToken?username=xxx&passwd=xxx
	//or http://127.0.0.1:8088/api/auth/token/getToken?client-token=xxx
	global.Router.GET("/api/auth/token/getToken", tokenAuth)
	//http://127.0.0.1:8088/api/auth/signin
	global.Router.GET("/api/auth/signin", signInAuth)
	global.Router.POST("/api/auth/signin", signInAuthP)

	global.Router.GET("/auth/signin", signInAuth)
	global.Router.POST("/auth/signin", signInAuthP)
	//
	go timeToClen()
}
