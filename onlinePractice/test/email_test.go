package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "测试人 <ranzhouol@163.com>"
	e.To = []string{"1844712096@qq.com"}
	e.Subject = "验证码发送测试"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("你的验证码是:<b>123456</b>")
	//err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "ranzhouol@163.com", "QFTPLLCJVEWAVEUX", "smtp.163.com"))
	// 返回 EOF 错误时，关闭SSL
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "ranzhouol@163.com", "QFTPLLCJVEWAVEUX", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
