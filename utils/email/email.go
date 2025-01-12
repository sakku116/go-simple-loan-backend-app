package email_util

import (
	"backend/config"
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(config config.EmailConfig, to []string, subject, body string) error {
	headers := make(map[string]string)
	headers["From"] = config.FromEmail
	headers["To"] = strings.Join(to, ",")
	headers["Subject"] = subject
	headers["Content-Type"] = "text/plain; charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", config.Username, config.Password, config.SMTPHost)

	addr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)
	err := smtp.SendMail(addr, auth, config.FromEmail, to, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
