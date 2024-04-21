package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstucture:"telegramToken"`
	GptToken      string `mapstucture:"gptToken"`
}

func LoadConfig(path string) (c *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
