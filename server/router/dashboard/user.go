package dashboard

import (
	"IOM/server/config"
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// getUserInfo 获取用户信息
// http://127.0.0.1:8088/api/dashboard/user/info
func getUserInfo(context *gin.Context) {
	logrus.Info("WebAPI: Request getUserInfo", context.Request.URL.Path)
	token := context.DefaultQuery("token", "")
	if !global.IsValidToken(token) && !global.IsCookieAuthValid(context) {
		context.String(http.StatusUnauthorized, "token is invalid")
		logrus.Errorln("WebAPI: Request ", context.Request.URL.Path, " token is invalid.")
		return
	}
	cookie, _ := context.Cookie("username")
	var userId int
	config.Db.QueryRow("select UserID from IOMSystem where Account=?", cookie).Scan(&userId)

	context.JSON(http.StatusOK, gin.H{
		"username":       cookie,
		"time":           0,
		"userId":         userId,
		"roleId":         0,
		"organizationId": 0,
		"postId":         0,
		"lastLoginTime":  config.StartTime.Format("2006-06-12 11:11:00"),
		"lastLoginIp":    context.ClientIP()})

}
