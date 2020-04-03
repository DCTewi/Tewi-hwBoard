package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
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
	mail.Content = `<body>
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
`
	return mail
}

// NewAdminMail template
func NewAdminMail(sendto []string, adminkey string) (mail Mail) {
	mail.Sendto = sendto
	mail.Nickname = config.App.Title
	mail.Subject = "[管理员登录]本次的登录链接"
	mail.Content = `<body>
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
`
	return mail
}

// Send mail attempt
func Send(mail Mail) error {
	now := time.Now()

	if last, ok := lastsent[mail.Sendto[0]]; ok {
		t, _ := time.ParseDuration("10m")
		if now.After(last.Add(t)) {
			err := trySendMail(mail)
			if err == nil {
				lastsent[mail.Sendto[0]] = now
			}
			return err
		}
		return errors.New("too many emails")
	}

	err := trySendMail(mail)
	if err == nil {
		lastsent[mail.Sendto[0]] = now
	}
	return err
}

func trySendMail(mail Mail) error {
	if config.Mail.SMTPPort != "25" {
		return sendDirectlyUsingTLS(mail)
	}
	return sendDirectly(mail)
}

func sendDirectly(mail Mail) error {
	account := &config.Mail
	auth := smtp.PlainAuth("", account.MailAccount, account.Password, account.SMTPServer)

	header := map[string]string{
		"From":         mail.Nickname + "<noreply@dctewi.com>",
		"To":           strings.Join(mail.Sendto, ","),
		"Subject":      mail.Subject,
		"Content-Type": "text/html; charset=UTF-8",
	}
	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	msg += "\r\n" + mail.Content

	return smtp.SendMail(account.SMTPServer+":"+account.SMTPPort, auth, account.MailAccount, mail.Sendto, []byte(msg))

}

func sendDirectlyUsingTLS(mail Mail) error {
	account := &config.Mail
	auth := smtp.PlainAuth("", account.MailAccount, account.Password, account.SMTPServer)
	header := map[string]string{
		"From":         mail.Nickname + "<noreply@dctewi.com>",
		"To":           strings.Join(mail.Sendto, ","),
		"Subject":      mail.Subject,
		"Content-Type": "text/html; charset=UTF-8",
	}
	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	msg += "\r\n" + mail.Content

	c, err := tlsDial(fmt.Sprintf("%s:%s", account.SMTPServer, account.SMTPPort))
	if err != nil {
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(account.MailAccount); err != nil {
		return err
	}
	for _, addr := range mail.Sendto {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func tlsDial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
