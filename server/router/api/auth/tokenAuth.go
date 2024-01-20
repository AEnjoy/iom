package auth

import (
	"IOM/server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tokenAuth(context *gin.Context) {
	username := context.DefaultQuery("username", "")
	passwd := context.DefaultQuery("passwd", "")
	var tokenApi *config.TokenApi
	if username != "" && passwd != "" {
		tokenApi.SetUserName(username)
		tokenApi.SetPassword(passwd)
		s, _ := tokenApi.GetTrustToken()
		tokenApis = append(tokenApis, tokenApi)
		context.String(http.StatusOK, s)
		return
	}
	clientToken := context.DefaultQuery("client-token", "")
	if clientToken != "" {
		tokenApi.SetClientToken(clientToken)
		s, _ := tokenApi.GetTrustToken()
		tokenApis = append(tokenApis, tokenApi)
		context.String(http.StatusOK, s)
		return
	}
	context.String(http.StatusUnauthorized, "not allow interface")

}
