# IOM -LoveO&M 爱运维

IOM是一套Server-Agent（管理面板与服务器-探针）的，对kubernetes MasterNode-ComputeNode或VMControlNode-VMComputerNode的可视化与自动化运维工具。

**正在努力的开发中**

## 功能支持：

> Server端（面板端）：
>
> 充当kubernetes控制端 kubctl的GUI界面，对连接上的计算节点支持进行自动化容器编排，创建容器。
>
> 充当虚拟机的管理界面，自动化管理虚拟机，批量生成虚拟机、为虚拟机添加密钥、执行命令、部署应用。
>
> 收集Agent上报的数据，提供数据监控大盘。
>
> *支持添加多个kubernetes集群，提供简单节点管理、命名空间管理，起到简化运维的目的*。
>
> 远程命令执行及网页ttyShell终端，对局域网内的管理使用标准shell连接，对跨网域的终端使用WebSocket连接。
>
> 软件包管理：安装、卸载、更新软件包
>
> 数据展示：
>
> 大屏列表展示尽可能多的节点状态，单节点数据展示及控制等

> Agent端（探针端）：
>
> 上报节点运行数据，执行面板端下发的命令

## 技术栈

Golang 1.20+，gRPC，Gin

Vue3 ElementPlus

## 已经实现的功能

1.查看数据监控大屏

# 构建说明

```
git clone https://github.com/aenjoy/iom

cd iom
cd dashboard
# 安装依赖
npm install
# 运行项目
npm run dev
# 打包发布
npm run build

# 编译服务端
cd ../server
# 安装依赖
go mod tidy 
# 运行项目
go build main.go

# 编译Agent端
cd ../agent
# 安装依赖
go mod tidy 
# 运行项目
go build main.go
```

