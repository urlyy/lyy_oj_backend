package util

import (
	"math/rand"

	"gopkg.in/gomail.v2"
)

const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func GenCaptcha(length int) string {
	b := make([]byte, length)
	for i := range b {
		idx := rand.Intn(len(charset)) //[0~len-1]
		b[i] = charset[idx]
	}
	captcha := string(b)
	return captcha
}

func SendEmail(to string, subject string, content string) error {
	smtp := &GetProjectConfig().SMTP
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", smtp.From)
	//接收人
	m.SetHeader("To", to)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "urlyy")
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", content)
	//附件
	//m.Attach("./myPic.png")
	// 第4个参数是填授权码
	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.Username, smtp.Password)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
