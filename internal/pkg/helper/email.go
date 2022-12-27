package helper

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/template"
	"go-start/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, subject string, tpl string, link string) error {
	fmt.Println(config.Cfg.Email.Host)
	fmt.Println(1)
	d := gomail.NewDialer(config.Cfg.Email.Host, int(config.Cfg.Email.Port), config.Cfg.Email.Address, config.Cfg.Email.Password)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(config.Cfg.Email.Address, "推书小站"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	body, err := template.ParseFiles(tpl)
	if err != nil {
		return err
	}
	preData := struct {
		VerifyLink string
	}{
		VerifyLink: link,
	}
	buf := new(bytes.Buffer)
	_ = body.Execute(buf, preData)
	m.SetBody("text/html", buf.String())
	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
