package auth

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
    GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
}

var AppConfig *Config

func InitConfig() {
    viper.SetConfigFile("/storage/emulated/0/rest_api/configs/config.env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    AppConfig = &Config{}
    if err := viper.Unmarshal(AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }
}

