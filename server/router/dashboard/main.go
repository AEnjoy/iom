package dashboard

import "IOM/server/global"

func Main() {
	global.Router.GET("/api/dashboard/user/info", getUserInfo)
}
