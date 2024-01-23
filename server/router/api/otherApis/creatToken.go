package otherApis

import (
	"IOM/server/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreatToken(context *gin.Context) {
	logrus.Info("WebAPI: Request ", context.Request.URL.Path)
	context.String(200, api.Capthca2())
}
