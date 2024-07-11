package load

import (
	"fmt"

	"github.com/spf13/viper"
)

type EmailConfig struct {
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUser     string `mapstructure:"SMTP_USER"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	EmailFrom    string `mapstructure:"EMAIL_FROM"`
}

func InitEmailConfig() (*EmailConfig, error) {
	viper.SetConfigFile("configs/email.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var emailConfig EmailConfig
	if err := viper.Unmarshal(&emailConfig); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &emailConfig, nil
}
