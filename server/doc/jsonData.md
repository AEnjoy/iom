# JSON 数据格式



## 设备数据JSON

请求API /api/device/:clientID(or token)/info
```json
{
	"ClientID": "",
	"HostState": {
		"CPU": 0,
		"MemUsed": 0,
		"SwapUsed": 0,
		"DiskUsed": 0,
		"NetInTransfer": 0,
		"NetOutTransfer": 0,
		"NetInSpeed": 0,
		"NetOutSpeed": 0,
		"Uptime": 0,
		"Load1": 0,
		"Load5": 0,
		"Load15": 0,
		"TcpConnCount": 0,
		"UdpConnCount": 0,
		"ProcessCount": 0
	},
	"HostInfo": {
		"Platform": "",
		"PlatformVersion": "",
		"CPU": null,
		"MemTotal": 0,
		"DiskTotal": 0,
		"SwapTotal": 0,
		"Arch": "",
		"Virtualization": "",
		"BootTime": 0,
		"CountryCode": "",
		"Version": ""
	},
	"DataTime": "0001-01-01T00:00:00Z",
    "FormatTime": "2024-1-19 23:35:02" 
}
```

### 请求设备在线设备API /api/devices/get-devices

由于在线的设备可以获取到系统信息，因此可以获取更加细致的Type内容 0：Windows 1Linux

请求所有设备/api/devices/get-all-devices(如果使用的是次请求，则Type为如下值0:StandDevices,1:PVE,2:OpenStack,3:k8sHost)

```json
[
  {
    "Token":"snlspjrnzxcevyfk",
    "Type":0
  }
]
```