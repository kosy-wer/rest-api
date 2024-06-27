package test

import (
    "github.com/spf13/viper"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestViper(t *testing.T) {
    var config *viper.Viper = viper.New()
    assert.NotNil(t, config)
}

func TestENV(t *testing.T) {
    config := viper.New()
    config.SetConfigFile("/storage/emulated/0/rest_api/configs/config.env")
    config.AutomaticEnv()

    // read config
    err := config.ReadInConfig()
    assert.Nil(t, err, "Failed to read config file: %v", err)

    assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
    assert.Equal(t, "Eko Kurniawan Khannedy", config.GetString("APP_AUTHOR"))
    assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
    assert.Equal(t, "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", config.GetString("GOOGLE_CLIENT_SECRET"))
    assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
    assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}

