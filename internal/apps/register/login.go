package register

import (
    "fmt"
    "github.com/spf13/viper"
)

func SayHello() {
    config := viper.New()
    config.SetConfigName("config")
    config.SetConfigType("json")
    config.AddConfigPath("/storage/emulated/0/rest_api/configs")
    

    // read config
    err := config.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config file:", err)
        return
    }

    username := config.GetString("app.name")
    fmt.Println("Hello", username, "from rest_api_login package!")
}

