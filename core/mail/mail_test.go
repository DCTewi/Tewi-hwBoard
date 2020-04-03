package mail_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dctewi/tewi-hwboard/core/mail"
)

func TestSendMail(t *testing.T) {
	err := mail.Send(mail.NewMail([]string{"dctewi@dctewi.com"}, "123456"))

	fmt.Println(err)
	err = mail.Send(mail.NewMail([]string{"dctewi@dctewi.com"}, "123456"))
	fmt.Println(err)
	err = mail.Send(mail.NewMail([]string{"dctewi@dctewi.com"}, "123456"))
	fmt.Println(err)

	t.Log("send mail test passed")
}

func TestSendAdminMail(t *testing.T) {
	err := mail.Send(mail.NewAdminMail([]string{"dctewi@dctewi.com"}, "32456789"))

	fmt.Println(err)
	err = mail.Send(mail.NewAdminMail([]string{"dctewi@dctewi.com"}, "32456789"))
	fmt.Println(err)
	err = mail.Send(mail.NewAdminMail([]string{"dctewi@dctewi.com"}, "32456789"))
	fmt.Println(err)

	t.Log("send mail test passed")
}

func TestTime(t *testing.T) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	ti, err := time.ParseInLocation("2006/01/02T15:04", "2020/02/02T04:22", l)
	fmt.Println(ti, err)
	t.Log("send mail test passed")
}
