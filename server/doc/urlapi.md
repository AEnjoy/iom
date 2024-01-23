Description of the route URL and parameters.

# /      根

# /api  后端API接口

所有/api下的请求均支持使用cookie和URL后跟随参数提供授权（/api/auth下**不**响应此参数）。请求参数为token=xxx

如curl  *http://127.0.0.1:8088/api/group/add?token=xxx&id=xxx&name=xxx*

不使用浏览器访问后端接口时，**token参数必须设置**

token/cookie不被授权时，将返回401 StatusUnauthorized。错误消息 “token is invalid”

请求参数错误时，将返回400 BadRequest或其它错误代码.

## /api/group模块 设备组API

### POST请求

#### /api/group/add 添加组

> 参数
>
> id: int设备组ID 
>
> name：string 组名 默认为groupName

成功返回200，ok

失败返回400，StatusBadRequest 请检查id是否重复(或值有效)

如果返回500，StatusInternalServerError 请检查后端数据库是否出现异常

### Delete请求

#### /api/group/delete

> 参数
>
> id：int 设备组ID 

成功返回200，ok

失败返回400，StatusBadRequest 请检查id是否存在(或值有效)

如果返回500，StatusInternalServerError 请检查后端数据库是否出现异常

## /api/auth模块 授权相关API

### GET/POST请求

#### /api/auth/signin 登录

> 参数
>
> username：string 用户名
>
> password：string 密码

成功返回302跳转/dashboard 且设置授权cookie

失败返回401 StatusUnauthorized 消息sign in failed. Username or password is not correct?



### GET请求

#### /api/auth/token/getToken 获取可信token

> 参数
>
> username：string 用户名
>
> passwd：string 密码
>
> 或
>
> client-token：string 通过已经生成的client密钥生成/获取token

根据所提供的用户名和密码获取可信的token

成功返回200，token

失败返回401 StatusUnauthorized

## /api/device模块 设备相关API

### POST请求

#### /api/devices/add 添加设备

> 参数
>
> deviceId：int(6) 6位设备ID
>
> devicesWeight: int value>=0 设备排序权重 默认为0
>
> devicesToken: string(16) 16位设备token 应该数字字母混合
>
> deviceName:string 设备名(不应超过64位字符)
>
> deviceFlag: int 设备标识 0:StandDevices,1:PVE,2:OpenStack,3:k8sHost 默认为0 
>
> groupId: int 设备归属组ID 将设备添至该组下 需要组存在

成功返回200，ok

失败返回500，错误信息 请检查值



#### /api/devices/getdevices获取所有**已上线**的设备列表

> 参数
>
> null

如果成功，返回200，JSON数据 设备列表 JSON定义请参考[jsonData.md](jsonData.md)



#### /api/devices/:ClientID/getpackages 获取设备已安装的程序包

或/api/devices/:id/getpackages 

> 参数
>
> id或clientID int/string

如果成功，返回200，JSON数据 设备列表 JSON定义请参考[jsonData.md](jsonData.md)



#### /api/devices/:ClientID/info 获取设备上报的数据

或/api/devices/:id/info

> 参数
>
> id或clientID int/string

如果成功，返回200，JSON数据 设备列表 JSON定义请参考[jsonData.md](jsonData.md)



### Delete请求

#### /api/devices/delete 删除设备

> 参数
>
> id：int 设备id
>
> groupId：int 设备所属组id

成功返回200，ok

失败返回500 StatusInternalServerError，请检查错误信息 设备id、组id是否正确

## /api/res模块 后端资源文件读取

### GET请求

#### /api/res/file/:file 获取后端脚本文件

> 参数
>
> :file 后端文件

成功返回200，脚本内容

失败返回404，找不到资源

## /api 其它api


### GET请求

#### /api/version 获取版本信息

> 参数
> 
>  null

成功返回200，IOM Server Version


#### /api/creat-token 创建随机token

> 参数
>
> null

成功返回200，16位数字字母混合 token

