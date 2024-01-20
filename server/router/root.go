package router

import (
	"IOM/server/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	global.Router.Static("/static", "./res")
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
