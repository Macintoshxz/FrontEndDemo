package test

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestTask(T *testing.T) {
	MailSendCode("1633516187@qq.com", "test-data")
}

func MailSendCode(emailstr, msg string) (err error) {

	e := email.NewEmail()
	e.From = "BAM <13609062201@163.com>"
	// xingyi1228@gmail.com
	e.To = []string{"xingyi1228@gmail.com", emailstr}
	e.Subject = "[Revisit PM] " + msg

	str := "This a reminder email: You marked 'pass & revisit' for [" + msg + "] in the pipeline database. It's now time to revisit it."
	e.HTML = []byte(str)
	err = e.SendWithTLS("smtp.163.com:465",
		// 注意这里的密码不是邮箱登录密码, 是开启 smtp 服务后获取的一串验证码
		smtp.PlainAuth("", "13609062201@163.com", "DSKSFGBDPPVVOKRV", "smtp.163.com"),
		// 指定跳过安全验证 ，
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		log.Println("发送邮件失败", err)
		panic(err)
	}
	return
}
