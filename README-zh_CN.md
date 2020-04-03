# Tewi-hwBoard

[For a English README](./README.md)

一个 Go 编写的提供发布、提交、收取等相关操作的作业管理系统Web App。

## 截屏

主页：
![](https://s1.ax1x.com/2020/04/03/GNbcge.png)

登录页：
![](https://s1.ax1x.com/2020/04/03/GNbgjH.png)

列表页：
![](https://s1.ax1x.com/2020/04/03/GNbRud.png)

管理员页：
![](https://s1.ax1x.com/2020/04/03/GNbfHI.png)

## 特性

- 部署简单
- 易于本地化
- 基于 Go 1.14.1
- 基于 SQLite3

## 部署

> 这是一种可能的部署的方法，测试于Ubuntu 18.04 LTS.

1. 配置环境

```shell
sudo apt install golang sqlite3 supervisor
go version
go env
sqlite3 -version
```

2. 获取本应用

```shell
go get github.com/dctewi/tewi-hwboard
```

3. 文件准备

进入想要安装的位置，如`~/hwboard`。将源码目录下的`app`文件夹复制进来，并在此处创建 SQLite3 数据库：

```shell
mkdir ~/hwboard ~/hwboard/app

cd ${GOPATH}/src/github.com/dctewi/tewi-hwboard/
cp ./app/* ~/hwboard/app/

cd ~/hwboard
sqlite3 database.db
sqlite> .read app/database.sql
sqlite> .quit
```

5. 编译并初次运行以生成配置文件

```shell
cd ~/hwboard
go build github.com/dctewi/tewi-hwboard
./tewi-hwboard
```

6. 应用设置

修改生成的配置文件`~/hwboard/config.json`的相关项目，该文件将在每次启动时读取。

7. 创建守护进程

编辑`/lib/systemd/system/hwboard.service`：

```ini
[Unit]
Description=hwboard
After=network.target

[Service]
User=yourusername
WorkingDirectory=/home/yourusername/hwboard/
ExecStart=/home/yourusername/hwboard/tewi-hwboard
Restart=always

[Install]
WantedBy=multi-user.target
```

然后通过`systemctl`启动服务：

```shell
sudo systemctl daemon-reload
sudo systemctl start hwboard.service
sudo systemctl status hwboard.service
```

到此，该应用已经安装完毕，访问设置的地址和端口即可访问。

## 联系

dcewi@dctewi.com

## 协议

![License-MIT](https://img.shields.io/badge/license-MIT-66ccff.svg)