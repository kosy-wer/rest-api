package load

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleLoginConfig  *oauth2.Config
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigFile("configs/googleOauth.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	AppConfig.GoogleLoginConfig = &oauth2.Config{
		ClientID:     AppConfig.GoogleClientID,
		ClientSecret: AppConfig.GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/google_callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
