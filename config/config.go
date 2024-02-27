package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	TelegramToken string `yaml:"telegramToken"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func init() {
	fileBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("error reading confiog file...")
	}

	ret := &Config{}
	err = yaml.Unmarshal(fileBytes, ret)
	if err != nil {
		log.Fatal("error getting config...")
	}
	config = *ret
}
