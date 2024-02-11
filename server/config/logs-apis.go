package config

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

type LoginLog struct {
	InfoId        int
	Username      string
	Ipaddr        string
	LoginLocation string
	Status        string //0:成功 1:失败
	Browser       string //ua
	Msg           string //Info
	CreateTime    time.Time
}

type LoginLogs struct {
	Total int
	Data  []LoginLog
}

func GetLogLogin(pageNum int, pageSize int, loginLocation string, username string) LoginLogs {
	var rows *sql.Rows
	var loginLogs LoginLogs
	loginLogs.Total = 0
	if username == "" {
		if loginLocation == "" {
			rows, _ = Db.Query(`SELECT * FROM LoginLog  limit ?`, pageSize)
		} else {
			rows, _ = Db.Query(`SELECT * FROM LoginLog where IP=? limit ?`, loginLocation, pageSize)
		}
	} else {
		if loginLocation == "" {
			rows, _ = Db.Query(`SELECT * FROM LoginLog where User=?  limit ?`, username, pageSize)
		} else {
			rows, _ = Db.Query(`SELECT * FROM LoginLog where User=? and IP=? limit ?`, username, loginLocation, pageSize)
		}
	}
	defer rows.Close()
	for rows.Next() {
		loginLogs.Total++
		var loginLog LoginLog
		var url string
		rows.Scan(&loginLog.InfoId, &loginLog.Username, &loginLog.Ipaddr, &url, &loginLog.Browser, &loginLog.Msg, &loginLog.CreateTime)
		loginLog.LoginLocation = url // todo 后面可以加IP地址查地理位置
		loginLog.Status = "0"
		loginLogs.Data = append(loginLogs.Data, loginLog)
	}
	return loginLogs
}
func AddLoginLog(username string, ipaddr string, url string, browser string, msg string, time time.Time) {
	logrus.Info("添加登录日志:username:", username, "ipaddr:", ipaddr, "url:", url, "browser:", browser, "msg:", msg, "time:", time)
	stmt, _ := Db.Prepare(`INSERT into LoginLog(User,IP,URL,UA,Info,Time) VALUES(?,?,?,?,?,?)`)
	stmt.Exec(username, ipaddr, url, browser, msg, time)
}
func DelLoginInfo(infoID int) error {
	logrus.Info("删除登录日志:infoID:", infoID)
	stmt, _ := Db.Prepare(`DELETE FROM LoginLog where ID=?`)
	_, err := stmt.Exec(infoID)
	logrus.Info("删除登录日志成功标志 为空则没有失败:", err)
	return err
}
func CleanAllLoginInfo() error {
	logrus.Info("清空所有登录日志")
	stmt, _ := Db.Prepare(`DELETE FROM LoginLog`)
	_, err := stmt.Exec()
	logrus.Info("清空所有登录日志成功标志 为空则没有失败:", err)
	return err
}

type OperLog struct {
	OperId       int       `json:"operId"`
	Title        string    `json:"title"`        //操作模块
	BusinessType string    `json:"businessType"` //操作类型
	Method       string    `json:"method"`
	OperName     string    `json:"operName"` //操作人员(url)
	OperUrl      string    `json:"operUrl"`
	OperIp       string    `json:"operIp"`
	OperLocation string    `json:"operLocation"` //操作地点
	Status       string    `json:"status"`
	Msg          string    `json:"errorMsg"`
	Time         time.Time `json:"createTime"`
}
type OperLogs struct {
	Total int `json:"total"`

	Data []OperLog `json:"data"`
}

func GetOperInfo(pageNo int, pageSize int, title string, operName string, businessType string) OperLogs {
	logrus.Info("获取操作日志:pageNo:", pageNo, "pageSize:", pageSize, "title:", title, "operName:", operName, "businessType:", businessType)
	var operLogs OperLogs
	var rows *sql.Rows
	rows, _ = Db.Query("SELECT * FROM WebOperaLog where Info like ? and Method like ? limit ?", "%"+title+"%", "%"+businessType+"%", pageSize)
	defer rows.Close()
	for rows.Next() {
		var operLog OperLog
		rows.Scan(&operLog.OperId, &operLog.Title, &operLog.Method, &operLog.OperName, &operLog.Time)
		operLog.OperUrl = operLog.OperName
		operLog.Status = "0"
		operLog.Msg = "null"
		operLog.OperIp = "127.0.0.1"
		operLogs.Data = append(operLogs.Data, operLog)
		operLogs.Total++
	}
	return operLogs
}
func AddOperLogs(title string, method string, url string) {
	logrus.Info("添加操作日志:title:", title, "method:", method, "url:", url)
	Db.Exec("INSERT INTO WebOperaLog(Info,Method,URL,Time) VALUES(?,?,?,?)", title, method, url, time.Now())
}
func DelOperLogs(id int) {
	logrus.Info("删除操作日志:id:", id)

	Db.Exec("DELETE FROM WebOperaLog WHERE ID=?", id)
}
func CleanAllOperLogs() {
	logrus.Info("清空操作日志")

	Db.Exec("DELETE FROM WebOperaLog")
}
