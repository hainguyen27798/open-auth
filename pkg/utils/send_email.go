package utils

import (
	"bytes"
	"fmt"
	"github.com/open-auth/global"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) string {
	msg := "MINE-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Name)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\n", mail.Body)
	return msg
}

func SendToEmail(templateName string, from string, to []string, data map[string]interface{}) error {
	body, err := getEmailTemplate(templateName, data)

	if err != nil {
		global.Logger.Error("Get email template error", zap.Error(err))
		return err
	}

	content := Mail{
		From:    EmailAddress{Address: from, Name: from},
		To:      to,
		Subject: "OTP Verification",
		Body:    body,
	}

	message := BuildMessage(content)

	SMTPHost := global.Config.SMTP.Host
	SMTPPort := global.Config.SMTP.Port
	SMTPUsername := global.Config.SMTP.Username
	SMTPPassword := global.Config.SMTP.Password

	// smtp auth
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	host := fmt.Sprintf("%s:%d", SMTPHost, SMTPPort)
	if err := smtp.SendMail(host, auth, from, to, []byte(message)); err != nil {
		fmt.Println(err.Error())
		global.Logger.Error("Send mail error:", zap.Error(err))
		return err
	}

	return nil
}

func getEmailTemplate(name string, data map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.ParseFiles(fmt.Sprintf("templates/%s.html", name)))
	err := t.Execute(htmlTemplate, data)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}
