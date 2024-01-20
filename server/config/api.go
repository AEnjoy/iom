package config

import (
	"errors"
	"fmt"
	"time"
)

func GroupAdd(groupId int, grupName string) error {
	_, err := Db.Exec("insert into GroupsA(ID,GroupName) values(?,?)", groupId, grupName)
	if err != nil {
		return err
	}
	var dt DevicesGroup
	dt.GroupName = grupName
	DevicesGroupEX[groupId] = dt
	return nil
}
func GroupNameGetGroupID(groupName string) int {
	rows, _ := Db.Query("SELECT ID FROM GroupsA where GroupName=?", groupName)
	defer rows.Close()
	for rows.Next() {
		var ID int
		rows.Scan(&ID)
		return ID
	}
	return 0
}

func GroupDelete(groupId int) error {
	_, err := Db.Exec("delete from GroupsA where ID = ?", groupId)
	if err != nil {
		return err
	}
	delete(DevicesGroupEX, groupId)
	return nil
}
func DeviceAdd(deviceId int, devicesWeight int, devicesToken string,
	deviceName string, deviceFlag int, groupId int) error {
	addedTime := time.Now()
	_, err := Db.Exec("insert into Devices(DevicesID,Weight,Token,DevicesName,flag,AddedTime) values(?,?,?,?,?,?)", deviceId, devicesWeight, devicesToken, deviceName, deviceFlag, addedTime)
	if err != nil {
		return err
	}
	_, err = Db.Exec("insert into DevicesGroup(GroupID,DevicesID) values (?,?)", groupId, deviceId)
	if err != nil {
		return err
	}
	var devices Devices
	devices.DevicesID = deviceId
	devices.DevicesName = deviceName
	devices.Weight = devicesWeight
	devices.Token = devicesToken
	devices.flag = deviceFlag
	devices.AddedTime = addedTime
	devices.LastDataTime = addedTime
	var dt DevicesGroup
	dt = DevicesGroupEX[groupId]
	dt.DevicesListID = append(dt.DevicesListID, devices)
	DevicesGroupEX[groupId] = dt
	return nil
}
func DeviceTokenGetDevicesID(Token string) int {
	var returns int
	Db.QueryRow("select DevicesID from Devices where Token=?", Token).Scan(&returns)
	return returns
}
func DeviceIDGetDeviceToken(id int) string {
	var returns string
	err := Db.QueryRow("select Token from Devices where DevicesID=?", id).Scan(&returns)
	if err != nil {
		return ""
	}
	return returns
}

func DeviceDelete(deviceId int, groupID int) error {
	_, err := Db.Exec("delete from Devices where DevicesID = ?", deviceId)
	if err != nil {
		return err
	}
	_, err = Db.Exec("delete from DevicesGroup where DevicesID = ?", deviceId)
	if err != nil {
		return err
	}
	var dt DevicesGroup
	dt = DevicesGroupEX[groupID]
	for i, v := range dt.DevicesListID {
		if v.DevicesID == deviceId {
			dt.DevicesListID = append(dt.DevicesListID[:i], dt.DevicesListID[i+1:]...)
			break
		}
	}
	return nil
}
func DevicesNameGetGroupName(deviceName string) (string, error) {
	for _, group := range DevicesGroupEX {
		for _, devices := range group.DevicesListID {
			if devices.DevicesName == deviceName {
				return group.GroupName, nil
			}
		}
	}
	return "", errors.New("找不到设备")
}
func DeviceIDGetGroupName(deviceID int) (string, error) {
	for _, group := range DevicesGroupEX {
		for _, devices := range group.DevicesListID {
			if devices.DevicesID == deviceID {
				return group.GroupName, nil
			}
		}
	}
	return "", errors.New("找不到设备")
}
func DeviceIDGetGroupID(deviceID int) (int, error) {
	for _, group := range DevicesGroupEX {
		for _, devices := range group.DevicesListID {
			if devices.DevicesID == deviceID {
				rows, _ := Db.Query("SELECT ID FROM GroupsA where GroupName = ?", group.GroupName)
				gID := 0
				//defer rows.Close()
				for rows.Next() {
					fmt.Scan(&gID)
				}
				return gID, nil
			}
		}
	}
	return 0, errors.New("找不到设备")
}
func DeviceTokenGetDeviceName(deviceToken string) (string, error) {
	var deviceName string
	err := Db.QueryRow("select  DevicesName from Devices where Token = ?", deviceToken).Scan(&deviceName)
	if err != nil {
		return "", err
	}
	return deviceName, nil
}

// ConfigEdit
// param ConfigID int 0添加，-1使用ConfigName搜索
func Edit(ConfigID int, ConfigName string, ConfigValue string) {
	if ConfigID == 0 {
		Add(ConfigName, ConfigValue)
		return
	}
	if ConfigID == -1 {
		Db.Exec("UPDATE Config SET ConfigValue = ? where ConfigName = ?", ConfigValue, ConfigName)
		return
	}
	Db.Exec("UPDATE Config SET ConfigName = ?,ConfigValue = ? WHERE ConfigID = ?", ConfigName, ConfigValue, ConfigID)
}

func Add(configName string, configValue string) {
	rows, _ := Db.Query("SELECT count(*) FROM Config")
	defer rows.Close()
	var id int
	rows.Next()
	rows.Scan(&id)
	Db.Exec("INSERT INTO Config(ConfigID,ConfigName,ConfigValue) VALUES(?,?,?)", id, configName, configValue)
}

func Get(ConfigID int) (string, string, error) {
	var ConfigName string
	var ConfigValue string
	Db.QueryRow("SELECT ConfigName,ConfigValue FROM Config WHERE ConfigID = ?", ConfigID).Scan(&ConfigName, &ConfigValue)
	return ConfigName, ConfigValue, nil
}
func AdminAccountEdit(AdminAccount string, AdminPassword string) error {
	_, err := Db.Exec("UPDATE IOMSystem SET Account = ?,Password = ? WHERE Account = ?", AdminAccount, AdminPassword, AdminAccount)
	if err != nil {
		return err
	}
	return nil
}

func AdminAccountAdd(AdminAccount string, AdminPassword string) error {
	_, err := Db.Exec("INSERT INTO IOMSystem(Account,Password) VALUES(?,?)", AdminAccount, AdminPassword)
	if err != nil {
		return err
	}
	return nil
}
func AdminAccountDel(AdminAccount string) error {
	_, err := Db.Exec("delete from IOMSystem where Account = ?", AdminAccount)
	if err != nil {
		return err
	}
	return nil
}
