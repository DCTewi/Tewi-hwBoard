# Tewi-hwBoard

[中文版本点击这里](./README-zh_CN.md)

A homework management web application in GoLang.

## Previews

Homepage：
![](https://s1.ax1x.com/2020/04/03/GNbcge.png)

LoginPage：
![](https://s1.ax1x.com/2020/04/03/GNbgjH.png)

ListPage：
![](https://s1.ax1x.com/2020/04/03/GNbRud.png)

AdminPage：
![](https://s1.ax1x.com/2020/04/03/GNbfHI.png)

## Features

- Easy to deploy
- Easy to localize
- Based on Go 1.14.1
- Based on SQLite3

## Deployment

> This is a template building procedure on Ubuntu 18.04 LTS.

1. Pre-install before deployment

```shell
sudo apt install golang sqlite3 supervisor
go version
go env
sqlite3 -version
```

2. Get this application

```shell
go get github.com/dctewi/tewi-hwboard
```

3. Create directory

Copy `./app/` into somewhere you want to run this app, such as `~/hwboard`, and create a sqlite3 database here:

```shell
mkdir ~/hwboard ~/hwboard/app

cd ${GOPATH}/src/github.com/dctewi/tewi-hwboard/
cp ./app/* ~/hwboard/app/

cd ~/hwboard
sqlite3 database.db
sqlite> .read app/database.sql
sqlite> .quit
```

5. Complie and run for first time and generate config.json

```shell
cd ~/hwboard
go build github.com/dctewi/tewi-hwboard
./tewi-hwboard
```

6. Setting up

Edit the generated `~/hwboard/config.json/` file to customize your app instace. The Json file will be reload when app launched.

7. Launch app by daemon

Create and edit `/lib/systemd/system/hwboard.service`：

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

Start service by `systemctl`:

```shell
sudo systemctl daemon-reload
sudo systemctl start hwboard.service
sudo systemctl status hwboard.service
```

Then you can access your board site on the domain:port you just set.

## Contact

dcewi@dctewi.com

## Lincense

![License-MIT](https://img.shields.io/badge/license-MIT-66ccff.svg)