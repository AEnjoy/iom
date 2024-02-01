package config

import (
	"IOM/server/global"
	"IOM/server/utils/debugTools"
	"IOM/server/utils/hashtool"
	"errors"
	"math"
	"math/rand"
	"strings"
	"time"
)

var TrustUsers []IOMSystem
var tokenApis []TokenApi

var StartTime = time.Now()
var startTimeStr = StartTime.Format("2006-01-02 15:04:05")
var onceTokenFunc = func() string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	n := 16
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
var onceToken = onceTokenFunc()

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
		t.setTrustToken(hashtool.MD5(t.clientToken + startTimeStr + onceToken))
		debugTools.PrintLogsOnlyInDebugMode("GetTrustToken ClientTokenMode: " + t.TrustToken + " " + t.clientToken + " " + startTimeStr + " " + onceToken)
		return t.TrustToken, nil
	}
	if t.password != "" && t.username != "" {
		t.setTrustToken(hashtool.MD5(t.username + t.password + startTimeStr + onceToken))
		debugTools.PrintLogsOnlyInDebugMode("GetTrustToken UserNameMode: " + t.TrustToken + " " + t.username + " " + t.password + " " + startTimeStr + " " + onceToken)
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
		global.TrustedTokens[t] = global.TokenTimeToLife{Valid: true, Time: time.Now()}
		tokenApis = append(tokenApis, tokenApi)
	}
	for s, _ := range global.DevicesInfos {
		var tokenApi TokenApi
		tokenApi.SetClientToken(s)
		t, _ := tokenApi.GetTrustToken()
		//println("Debug: TrustToken is:", t)
		global.TrustedTokens[t] = global.TokenTimeToLife{Valid: true, Time: time.Now()}
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

func SaveOnceToken(token string) error {
	var retVal error
	_, retVal = Db.Exec("insert into OnceToken (Token,Time) values (?,?)", token, time.Now())
	return retVal
}

func SaveOnceToken2() error {
	var retVal error
	for v := range global.TrustedTokens {
		retVal = SaveOnceToken(v)
	}
	return retVal
}

func DeleteOnceToken(token string) error {
	var retVal error
	_, retVal = Db.Exec("delete from OnceToken where Token=?", token)
	return retVal
}

func LoadOnceToken() error {
	var retVal error
	rows, retVal := Db.Query("select * from OnceToken")
	if retVal != nil {
		return retVal
	}
	for rows.Next() {
		var t time.Time
		var token string
		retVal = rows.Scan(&token, &t)
		if retVal != nil {
			return retVal
		}
		if math.Abs(float64(t.Unix()-StartTime.Unix())) > 60*60*24 {
			retVal = DeleteOnceToken(token)
			if retVal != nil {
				return retVal
			}
		}
		//没有过期
		global.TrustedTokens[token] = global.TokenTimeToLife{Valid: true, Time: t}
	}
	return retVal
}
