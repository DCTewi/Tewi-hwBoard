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

3. Fill app settings

Edit the  `App`，`Mail` and `Database` settings in `${GOPATH}/src/github.com/dctewi/tewi-hwboard/config/config.go`. Setting others if you need.

```go
// App configs
var App = struct {
	Domain       string
	HTTPPort     string
	SSLPort      string
	UseTLS       bool
	TLSCrtPath   string
	TLSKeyPath   string
	Title        string
	UserTimeZone string
	AdminEmails  []string
}{
	Domain:       "localhost",
	HTTPPort:     ":80",
	SSLPort:      ":443",
	UseTLS:       true,
	TLSCrtPath:   "./domain.crt",
	TLSKeyPath:   "./domain.key",
	Title:        "作业布告栏",
	UserTimeZone: "Asia/Shanghai",
	AdminEmails: []string{
		"you@domain.com",
	},
}

// Mail configs
var Mail = struct {
	MailAccount string
	Password    string
	SMTPServer  string
	SMTPPort    string
}{
	MailAccount: "you@domain.com",
	Password:    "password",
	SMTPServer:  "smtp.somedomain.com",
	SMTPPort:    "465",
}

// Database configs
var Database = struct {
	Path string
}{
	Path: "./database.db",
}
```

4. Create directory

Copy `./config/database.sql`，`./static/*`，`./views/*` into somewhere you want to run this app, such as `~/hwboard`, and create a sqlite3 database here:

```shell
mkdir ~/hwboard ~/hwboard/static ~/hwboard/views

cd ${GOPATH}/src/github.com/dctewi/tewi-hwboard/
cp ./config/database.sql ~/hwboard/database.sql
cp ./static/* ~/hwboard/static/
cp ./views/* ~/hwboard/views/

cd ~/hwboard
sqlite3 database.db
sqlite> .read database.sql
sqlite> .quit
```

5. Complie 

```shell
cd ~/hwboard
go build github.com/dctewi/tewi-hwboard
```

6. Launch app by daemon

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