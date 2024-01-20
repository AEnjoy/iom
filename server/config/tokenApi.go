package config

import (
	"IOM/server/global"
	"IOM/server/utils/hashtool"
	"errors"
)

var TrustUsers []IOMSystem
var tokenApis []TokenApi

type TokenApi struct {
	username string
	password string
	//or
	clientToken string
	TrustToken  string
}

func (t *TokenApi) SetUserName(username string) {
	t.username = username
}
func (t *TokenApi) SetPassword(password string) {
	t.password = password
}
func (t *TokenApi) SetClientToken(clientToken string) {
	t.clientToken = clientToken
}
func (t *TokenApi) setTrustToken(TrustToken string) {
	t.TrustToken = TrustToken
}

func (t *TokenApi) GetTrustToken() (string, error) {
	if t.clientToken != "" {
		t.setTrustToken(hashtool.MD5(t.clientToken))
		return t.TrustToken, nil
	}
	if t.password != "" && t.username != "" {
		t.setTrustToken(hashtool.MD5(t.username + t.password))
		return t.TrustToken, nil
	}
	return "", errors.New("no trust token")
}

func tokenInit() {
	for _, user := range TrustUsers {
		var tokenApi TokenApi
		//println("Debug: Ac and Pw is:", user.Account, user.Password)
		tokenApi.SetPassword(user.Password)
		tokenApi.SetUserName(user.Account)
		t, _ := tokenApi.GetTrustToken()
		//println("Debug: TrustToken is:", t)
		global.TrustedTokens[t] = true
		tokenApis = append(tokenApis, tokenApi)
	}
	for s, _ := range global.DevicesInfos {
		var tokenApi TokenApi
		tokenApi.SetClientToken(s)
		t, _ := tokenApi.GetTrustToken()
		//println("Debug: TrustToken is:", t)
		global.TrustedTokens[t] = true
		tokenApis = append(tokenApis, tokenApi)
	}
}

func IsUserInTrustUsers(in IOMSystem) bool {
	var r = false
	for _, user := range TrustUsers {
		if user.Account == in.Account &&
			user.Password == in.Password {
			r = true
			break
		}
	}
	return r
}
