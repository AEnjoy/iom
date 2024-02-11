package router

import (
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func static() {
	global.Router.Static("/static", "./res")
	global.Router.Static("/assets", "./res")
	global.Router.Static("/index.html", "./index.html")
	global.Router.GET("/favicon.ico", func(context *gin.Context) {
		context.Redirect(302, "/static/favicon.ico")
	})
	global.Router.LoadHTMLGlob("res/html/*.html")
	global.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "404",
		})
	})
}
