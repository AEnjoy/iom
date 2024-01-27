package config

import (
	"IOM/server/global"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func DataBaseInit() {
	err := os.Mkdir("./data", os.ModePerm)
	if err != nil {
		logrus.Error("MakeDir:./data fail. Maybe exist?")
		//return
	}
	db, err := sql.Open("sqlite3", "file:data/config.db?mode=rwc")
	if err != nil {
		logrus.Errorln("DataBase:Init fail.", err)
	}
	logrus.Info("Init DataBase.")
	db.Exec("CREATE TABLE IOMSystem(Account char(12), Password CHAR(64))")
	db.Exec("INSERT INTO IOMSystem(Account,Password) VALUES('admin', 'passwords')")
	db.Exec("CREATE TABLE Devices(DevicesID int primary key,Weight int,Token char(18),DevicesName char(64),flag int,AddedTime time,LastDataTime time null)")
	db.Exec("CREATE TABLE GroupsA(ID int primary key, GroupName CHAR(64))")
	db.Exec("CREATE TABLE DevicesGroup(GroupID int references GroupsA(ID),DevicesID int references Devices(DevicesID));")
	db.Exec("create table Config(ConfigID int primary key, ConfigName char(64), ConfigValue char(64))")
	db.Exec("insert into Config(ConfigID,ConfigName,ConfigValue) values(1,'Port', '10000')")
	db.Exec("insert into Config(ConfigID,ConfigName,ConfigValue) values(2,'WebPort', '8088')")
	Db = db
	logrus.Info("Done.")
}
func DBOpen() {
	logrus.Info("Read Database")
	_, err := os.Stat("data/config.db")
	if err == nil {
		Db, err = sql.Open("sqlite3", "file:data/config.db?mode=rwc")
		if err != nil {
			logrus.Errorln("DataBase:Open fail.", err)
			os.Exit(137)
		}

	} else {
		if os.IsNotExist(err) {
			logrus.Errorln("DataBase:Open fail:DataBase config file are not exist.")
			DataBaseInit()
			//DBOpen()
		}
	}
	/*
		rows, _ := Db.Query("SELECT * FROM IOMSystem")
		for rows.Next(){
			var systemInfo IOMSystem
			rows.Scan(&systemInfo.Account, &systemInfo.Password)
			SystemInfo = append(SystemInfo,  systemInfo)
		}
		rows.Close()*/
	rows, e := Db.Query("SELECT * FROM Config")
	if e != nil {
		println(e.Error())
	}
	for rows.Next() {
		var ID int
		var configName string
		var configValue string
		rows.Scan(&ID, &configName, &configValue)
		switch ID {
		case 1:
			global.RPCPort = configValue
		case 2:
			global.WebPort = configValue
		}
	}
	rows, _ = Db.Query("select * from IOMSystem")
	for rows.Next() {
		var tSystemInfo IOMSystem
		rows.Scan(&tSystemInfo.Account, &tSystemInfo.Password)
		TrustUsers = append(TrustUsers, tSystemInfo)
	}
	rows.Close()
	rows, _ = Db.Query("SELECT * FROM GroupsA")
	defer rows.Close()
	for rows.Next() {
		var dt DevicesGroup
		var ID int
		err = rows.Scan(&ID, &dt.GroupName)
		rowst, _ := Db.Query("SELECT DevicesID FROM DevicesGroup where GroupID=?", strconv.Itoa(ID))
		for rowst.Next() {
			//println("Debug 86")
			var dt1 Devices
			err = rowst.Scan(&dt1.DevicesID)
			rowst1, _ := Db.Query("SELECT * FROM Devices where DevicesID=?", dt1.DevicesID)
			rowst1.Next()
			rowst1.Scan(&dt1.DevicesID, &dt1.Weight, &dt1.Token, &dt1.DevicesName, &dt1.Flag, &dt1.AddedTime, &dt1.LastDataTime)
			if dt1.Token == "" {
				continue
			}
			dt.DevicesListID = append(dt.DevicesListID, dt1)
			var dt3 global.DevicesInfo
			dt3.Online = false
			dt3.ClientID = dt1.Token
			global.DevicesInfos[dt1.Token] = dt3
			//println("Debug:100", global.DevicesInfos[dt1.Token].ClientID)
			rowst1.Close()
		}
		rowst.Close()
		DevicesGroupEX[ID] = dt
	}
	tokenInit()
}
func DBExitSave() {
	defer Db.Close()
	for token, info := range global.DevicesInfos {
		//其实随便什么都可以，只要标识这个设备存在就行
		if info.HostState.MemUsed != 0 {
			logrus.Info("DataBase:Save LastDataTime to device. Device token is:" + token)
			Db.Exec("UPDATE Devices SET LastDataTime=? WHERE Token=?", info.DataTime, token)
		}
	}
}
