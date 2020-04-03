package mail

import (
	"errors"
	"net/smtp"
	"strings"
	"time"

	"github.com/dctewi/tewi-hwboard/config"
)

// Mail object
type Mail struct {
	Sendto   []string
	Nickname string
	Subject  string
	Content  string
}

var lastsent = make(map[string]time.Time)

// NewMail template
func NewMail(sendto []string, confirmkey string) (mail Mail) {
	mail.Sendto = sendto
	mail.Nickname = config.App.Title
	mail.Subject = "[验证码] 来自 " + config.App.Title + " 的验证码"
	mail.Content = `<html>
<head>
<title>你的验证码是: ` + confirmkey + `</title>
</head>
<body>
<div style="background-color:#ECECEC; padding: 35px;">
<table cellpadding="0" align="center"
	   style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
	<tbody>
	<tr>
		<th valign="middle"
			style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #42a3d3; background-color: #49bcff; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
			<font face="微软雅黑" size="5" style="color: rgb(255, 255, 255); ">[验证码] 来自` + config.App.Title + `的验证码</font>
		</th>
	</tr>
	<tr>
		<td>
			<div style="padding:25px 35px 40px; background-color:#fff;">
				<p>你的验证码是: <br>
					<h1>` + confirmkey + `</h1>
					在10分钟内有效<br></p>
				<p align="right">` + config.App.Title + `</p>
				<p align="right">` + time.Now().Format("2006/1/2 15:04:05") + `</p>
				<div style="width:700px;margin:0 auto;">
					<div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
						<p>此邮件来自 ` + config.App.Domain + ` 的登录事务，若并非您本人操作, 请忽略此邮件.</p>
						<p>©dctewi@dctewi.com</p>
					</div>
				</div>
			</div>
		</td>
	</tr>
	</tbody>
</table>
</div>
</body>
</html>
`
	return mail
}

// NewAdminMail template
func NewAdminMail(sendto []string, adminkey string) (mail Mail) {
	mail.Sendto = sendto
	mail.Nickname = config.App.Title
	mail.Subject = "[管理员登录]本次的登录链接"
	mail.Content = `<html>
<head>
<title>管理员登录</title>
</head>
<body>
<div style="background-color:#ECECEC; padding: 35px;">
<table cellpadding="0" align="center"
	   style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
	<tbody>
	<tr>
		<th valign="middle"
			style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #42a3d3; background-color: #49bcff; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
			<font face="微软雅黑" size="5" style="color: rgb(255, 255, 255); ">[管理员登录]本次的登录链接</font>
		</th>
	</tr>
	<tr>
		<td>
			<div style="padding:25px 35px 40px; background-color:#fff;">
				<p>本次的登录链接为: <br>
					<a target="_blank" href="` + config.App.Domain + `/admin?akey=` + adminkey + `">` + config.App.Domain + `/admin?akey=` + adminkey + `</a><br />
					在10分钟内有效<br></p>
				<p align="right">` + config.App.Title + `</p>
				<p align="right">` + time.Now().Format("2006/1/2 15:04:05") + `</p>
				<div style="width:700px;margin:0 auto;">
					<div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
						<p>此邮件来自 ` + config.App.Domain + ` 的登录事务，若并非您本人操作, 请忽略此邮件.</p>
						<p>©dctewi@dctewi.com</p>
					</div>
				</div>
			</div>
		</td>
	</tr>
	</tbody>
</table>
</div>
</body>
</html>
`
	return mail
}

// Send mail attempt
func Send(mail Mail) error {
	now := time.Now()

	if last, ok := lastsent[mail.Sendto[0]]; ok {
		t, _ := time.ParseDuration("10m")
		if now.After(last.Add(t)) {
			err := sendDirectly(mail)
			if err == nil {
				lastsent[mail.Sendto[0]] = now
			}
			return err
		}
		return errors.New("too many emails")
	}

	err := sendDirectly(mail)
	if err == nil {
		lastsent[mail.Sendto[0]] = now
	}
	return err
}

func sendDirectly(mail Mail) error {
	account := &config.Mail
	auth := smtp.PlainAuth("", account.MailAccount, account.Password, account.SMTPServer)
	contentType := "Content-Type: text/html; charset=UTF-8"

	msg := []byte("To: " + strings.Join(mail.Sendto, ",") + "\r\nFrom: " + mail.Nickname +
		"<" + account.MailAccount + ">\r\nSubject: " + mail.Subject + "\r\n" + contentType + "\r\n\r\n" + mail.Content)

	return smtp.SendMail(account.SMTPServer+":"+account.SMTPPort, auth, account.MailAccount, mail.Sendto, msg)

}
