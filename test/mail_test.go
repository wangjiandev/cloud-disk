package test

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"
	"wangjiandev/cloud-disk/core/define"

	"github.com/jordan-wright/email"
)

func TestSendMall(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <wangjiantx2@163.com>"
	e.To = []string{"928647866@qq.com"}
	e.Subject = "验证码"
	e.HTML = []byte("<h1>您的验证码: 876589 </h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "wangjiantx2@163.com", define.MailKey, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("发送成功!")

}
