package otherApis

import (
	"IOM/server/api"
	"IOM/server/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreatToken(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	config.AddOperLogs("WebAPI访问：CreatToken", "GET", context.Request.URL.Path)
	context.String(200, api.Capthca2())
}
