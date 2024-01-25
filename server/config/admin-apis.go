package config

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
