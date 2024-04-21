package main

import (
	"github.com/iamvikaskumar/telegram-gpt/config"
	"github.com/iamvikaskumar/telegram-gpt/pkg/gpt"
	"github.com/iamvikaskumar/telegram-gpt/pkg/telegram"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		panic("error loading config...")
	}

	gptClient := gpt.GetClient(config.GptToken)

	if err != nil {
		panic("error getting response from chat gpt")
	}

	telegramClient := telegram.NewClient(config.TelegramToken, gptClient)
	telegramClient.Listen()
}
