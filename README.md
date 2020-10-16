## 安装
1. 下载
```shell script
git clone https://github.com/kudoochui/kudosServer.git
```
2. 启动注册中心
[安装consul](https://learn.hashicorp.com/consul/getting-started/install)
```
consul agent --dev
```
3. 运行游戏
```shell script
go build app/main.go
./main
```
   
## 分布式部署
```shell script
//启动一个连接服
./main -type gate -id gate-1
//再启动一个
./main -type gate -id gate-2