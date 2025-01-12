package config

type EmailConfig struct {
	SMTPHost  string
	SMTPPort  int
	Username  string
	Password  string
	FromEmail string
}

func NewGMailConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost:  Envs.GMAIL_SMTP_HOST,
		SMTPPort:  Envs.GMAIL_SMTP_PORT,
		Username:  Envs.GMAIL_USERNAME,
		Password:  Envs.GMAIL_PASSWORD,
		FromEmail: Envs.GMAIL_FROM_EMAIL,
	}
}
