package envTool

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ProxmoxEnv struct {
	username string
	password string
}

func (e ProxmoxEnv) ApiInit() {
	url := "https://127.0.0.1:8006/api2/json/access/ticket"
	info := make(map[string]string)
	info["username"] = e.username
	info["password"] = e.password
	bytesData, err := json.Marshal(info)
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		logrus.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		logrus.Error(err)
		return
	}

	//fmt.Println(response)
}
func (e ProxmoxEnv) SetUserName(username string) {
	e.username = username
}
func (e ProxmoxEnv) SetPassWord(password string) {
	e.password = password
}
